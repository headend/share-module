package configuration

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"github.com/headend/share-module/configuration/static-config"
	"github.com/headend/share-module/file-and-directory"
)

//GlobalConf lay conf Global
var GlobalConf *Conf

func (c *Conf) LoadConf() *Conf {
	fileName := static_config.ConfigFilePath
	if c.ConfigureFile != "" {
		fileName = c.ConfigureFile
	}
	confFile := file_and_directory.MyFile{Path:fileName}
	yamlFile, err := confFile.Read()
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
		panic(err)
	}
	err = yaml.Unmarshal([]byte(yamlFile), c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

//LoadConfFrom load conf from filename(para)
func LoadConfFromFile(filename string) error {
	c := &Conf{}
	if yamlFile, err := ioutil.ReadFile(filename); err == nil {
		if err := yaml.Unmarshal(yamlFile, c); err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}

