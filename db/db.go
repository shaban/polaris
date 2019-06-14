package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/lib/pq" //postgres
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

//EveDB is the entirety of the eve database tables
//implemented as maps of database id to actual type
type EveDB struct {
	Mapping map[string]mapper
}

func (e *EveDB) addMapToMapping(m mapper) {
	name := m.TableName()
	e.Mapping[name] = m
}

type mapper interface {
	GetByKey(int) interface{}
	SaveToDB() error
	FileName() string
	LoadFromDB() error
	TableName() string
	New(int, []byte) error
	LoadFromYAML() error
}

var (
	Blueprints   = blueprints(make(map[int]*Blueprint))
	CategoryIDs  = categoryIDs(make(map[int]*CategoryID))
	Certificates = certificates(make(map[int]*Certificate))
	GraphicIDs   = graphicIDs(make(map[int]*GraphicID))
	GroupIDs     = groupIDs(make(map[int]*GroupID))
	IconIDs      = iconIDs(make(map[int]*IconID))
	MarketGroups = marketGroups(make(map[int]*MarketGroup))
	Skins        = skins(make(map[int]*Skin))
	TypeIDs      = typeIDs(make(map[int]*TypeID))

	connectionString    string
	basePath            string
	yamlRelationalPath  string
	yamlObjectPath      string
	yamlExt             string
	yamlRelationalFiles []string
	yamlObjectFiles     []string

	pg *sql.DB
	//eve is an instance of eveDB
	eve *EveDB
)

//Open reads the configuration opens the postgres database
//sets path variables and boostraps the type mapper interface
//finally it loads all static data on an audit from yaml
//otherwise from database into the mapper map
func Open(conf *viper.Viper) {
	var (
		err error
	)
	connectionString = fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s",
		conf.GetString("database.postgres.user"),
		conf.GetString("database.postgres.db"),
		conf.GetString("database.postgres.password"),
		conf.GetString("database.postgres.host"),
	)

	if err = open(); err != nil {
		log.Fatal(err)
	}

	basePath = conf.GetString("path")
	yamlRelationalPath = conf.GetString("database.yaml.relational.path")
	yamlObjectPath = conf.GetString("database.yaml.object.path")
	yamlExt = conf.GetString("database.yaml.extension")
	yamlRelationalFiles = conf.GetStringSlice("database.yaml.relational.files")

	eve = new(EveDB)
	eve.Mapping = make(map[string]mapper)

	eve.addMapToMapping(Blueprints)
	eve.addMapToMapping(CategoryIDs)
	eve.addMapToMapping(Certificates)
	eve.addMapToMapping(GraphicIDs)
	eve.addMapToMapping(GroupIDs)
	eve.addMapToMapping(IconIDs)
	eve.addMapToMapping(MarketGroups)
	eve.addMapToMapping(Skins)
	eve.addMapToMapping(TypeIDs)

	for _, mappedGroup := range eve.Mapping {
		if tableExists(mappedGroup.TableName()) {
			//then load database instead
			if err = mappedGroup.LoadFromDB(); err != nil {
				log.Fatalf("Can't Load %s From Database: %s", mappedGroup.TableName(), err)
			}
			continue
		}
		if err = mappedGroup.LoadFromYAML(); err != nil {
			log.Fatalf("Can't Load %s From Yaml: %s", mappedGroup.FileName(), err)
		}
		createIfNotExists(mappedGroup.TableName())
		if err = mappedGroup.SaveToDB(); err != nil {
			log.Fatalf("Can't Save %s To Database: %s", mappedGroup.TableName(), err)
		}
	}
}

//EncodeByTableAndKey accesses items by tablename, row id
//and then encodes the row in JSON into a writer interface
func EncodeByTableAndKey(w io.Writer, table string, key int) error {
	var (
		v interface{}
	)
	if v = eve.Mapping[table].GetByKey(key); v == nil {
		return fmt.Errorf("Can't get table:%s key:%v is nil", table, key)
	}
	enc := json.NewEncoder(w)
	return enc.Encode(v)
}

//Init opens the database and stores it as a private
//package variable
func open() error {
	var err error
	pg, err = sql.Open("postgres", connectionString)
	return err

}

//Close the database all errors will make
//the Program crash so we don't have
//to deal with errors on defer
func Close() {
	if err := pg.Close(); err != nil {
		log.Fatalf("Can't Close Database: %s", err)
	}
}

//CreateIfNotExists creates a table as a key jsonb store
func createIfNotExists(table string) {
	if _, err := pg.Exec(`CREATE TABLE IF NOT EXISTS ` + table + ` (id integer PRIMARY KEY, data jsonb NOT NULL)`); err != nil {
		log.Fatalf("Can't create Table %s: %s", table, err)
	}
}

//doesTableExist returns wether the table exists
func tableExists(table string) bool {
	var exists = new(sql.NullString)
	row := pg.QueryRow(fmt.Sprintf("SELECT to_regclass('public.%s')", table))
	if err := row.Scan(exists); err != nil {
		log.Fatalf("Can't check if Table exists: %s", err)
	}
	return exists.Valid
}

//Insert a datastructure as JSONB into the database
func insert(table string, key int, data interface{}) error {
	var (
		err      error
		jsonData []byte
	)
	if jsonData, err = json.Marshal(data); err != nil {
		return err
	}
	_, err = pg.Exec(`INSERT INTO `+table+`(id, data) VALUES($1, $2)`, key, jsonData)
	return err
}

func loadFromDB(tt mapper) error {
	var (
		rows *sql.Rows
		err  error
		id   int
		data []byte
	)
	rows, err = pg.Query(fmt.Sprintf("SELECT * FROM %s", tt.FileName()))
	for rows.Next() {
		if err = rows.Scan(&id, &data); err != nil {
			return fmt.Errorf("Can't read value from %s: %s", tt.FileName(), err)
		}

		if err = tt.New(id, data); err != nil {
			return err
		}
	}
	return nil
}

//LoadYaml decodes a yaml file into a suitable data structure
//on specified yamlPath filepath
func loadYaml(dst interface{}, yamlPath string) error {
	var (
		err error
		f   *os.File
		dec *yaml.Decoder
	)
	// open yaml file
	if f, err = os.OpenFile(yamlPath, os.O_RDONLY, 0644); err != nil {
		return err
	}
	//decode into our destination in strict mode
	dec = yaml.NewDecoder(f)
	dec.SetStrict(true)
	defer f.Close()
	if err = dec.Decode(dst); err != nil {
		return err
	}
	log.Println(yamlPath, "loaded from YAML")
	return nil
}
