package parser

import (
	"io/ioutil"
	"path/filepath"

	"github.com/go-jack/util"
	"gopkg.in/yaml.v2"
)

//Config is the struct for our config file we're parsing
type Config struct {
	VerifyToken string `yaml:"verify_token"`
	AccessToken string `yaml:"access_token"`
	AppSecret   string `yaml:"app_secret"`
}

//Err is a simple Error handler/logger that logs error and reason (what I did) to probably lead to the error
var Err = util.HandleErr

//ReadYml is a method that allows us Read and parse the content of the File
func (x *Config) ReadYml() *Config {
	configFile, err := filepath.Abs("./bot.config.yml")

	if err != nil {
		Err(err, "TRYING TO RELATIVELY OPEN FILE PATH")
	}

	yamlFile, err := ioutil.ReadFile(configFile)

	if err != nil {
		Err(err, "TRYING TO READ CONFIG FILE")
	}

	ymlError := yaml.Unmarshal(yamlFile, &x)

	if ymlError != nil {
		Err(err, "TRYING TO UNMARSHALL THE CONFIG FILE")
	}

	return x
}
