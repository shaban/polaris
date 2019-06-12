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
		fileName string
		f   *os.File
		dec *yaml.Decoder
	)
	for _, fileName = range yamlFiles {
		//is data already loaded?
		loadable, needsCustomLoader := eve.Mapping[fileName].(customLoader)
		if tableExists(fileName) {
			//then load database instead
			goto loadFromDatabase
		}
		//if no load yaml files
		filePath = fmt.Sprintf("%s/%s/%s.%s", basePath, yamlPath, fileName, yamlExt)
		if needsCustomLoader {
			if f, err = os.OpenFile(filePath, os.O_RDONLY, 0644); err != nil {
				return err
			}
			//decode into our destination in strict mode
			dec = yaml.NewDecoder(f)
			dec.SetStrict(true)
			defer f.Close()
			if err = loadable.LoadFromYAML(dec);err!=nil{
				return err
			}
			goto createTable
		}
		if err = loadYaml(eve.Mapping[fileName], filePath); err != nil {
			log.Fatalf("Can't load Yaml %s in %s :%s", fileName, filePath, err)
		}
		createTable:
		//create a table from the extensionless filename
		createIfNotExists(fileName)
		if err = eve.Mapping[fileName].SaveToDB(); err != nil {
			return err
		}
		continue
	loadFromDatabase:
		if err = eve.Mapping[fileName].LoadFromDB(); err != nil {
			return err
		}
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
