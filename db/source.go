package db

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func loadStaticData() error {
	var (
		filePath string
		err      error
		keys     []int
		fileName string
	)

	//if no load yaml files
	for _, fileName = range yamlFiles {
		//is data already loaded?
		if tableExists(fileName) {
			//then load database instead
			goto loadFromDatabase
		}
		filePath = fmt.Sprintf("%s/%s/%s.%s", basePath, yamlPath, fileName, yamlExt)
		if err = loadYaml(eve.Mapping[fileName], filePath); err != nil {
			log.Fatalf("Can't load Yaml %s in %s", fileName, filePath)
		}
		//create a table from the extensionless filename
		createIfNotExists(fileName)
		keys = eve.keys(fileName)
		for _, key := range keys {
			row, ok := eve.byFileName(fileName, key)
			if !ok {
				return fmt.Errorf("Inserting Rows: Key %v does not exist in %s", key, fileName)
			}

			if err = Insert(fileName, key, row); err != nil {
				log.Fatalf("Can't insert %s: %s", fileName, err)
			}

		}
		return nil
	loadFromDatabase: 
	loadTable(fileName)
	}
	return err
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
