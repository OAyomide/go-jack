package parser

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

//Config is the struct for our config file we're parsing
type Config struct {
	VerifyToken string `yaml:"verify_token"`
	AccessToken string `yaml:"access_token"`
	AppSecret   string `yaml:"app_secret"`
}

//ReadYml is a method that allows us Read and parse the content of the File
func (x *Config) ReadYml() *Config {
	configFile, err := filepath.Abs("./bot.config.yml")

	if err != nil {
		panic(err, "TRYING TO RELATIVELY OPEN FILE PATH")
	}

	yamlFile, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(err, "TRYING TO READ CONFIG FILE")
	}

	ymlError := yaml.Unmarshal(yamlFile, &x)

	if ymlError != nil {
		panic(err, "TRYING TO UNMARSHALL THE CONFIG FILE")
	}

	return x
}
