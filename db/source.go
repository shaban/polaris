package db

import (
	"os"

	"gopkg.in/yaml.v2"
)

//LoadYaml decodes a yaml file into a suitable data structure
//on specified yamlPath filepath
func LoadYaml(dst interface{}, yamlPath string) error {
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
	return dec.Decode(dst)
}
