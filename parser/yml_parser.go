package parser

import (
	"io/ioutil"
	"path/filepath"

	"github.com/messenger"
	"gopkg.in/yaml.v2"
)

//Config is the struct for our config file we're parsing
type Config struct {
	VerifyToken    string                        `yaml:"verify_token"`
	AccessToken    string                        `yaml:"access_token"`
	AppSecret      string                        `yaml:"app_secret"`
	PersistentMenu []messenger.CallToActionsItem `yaml:"persistent_menu"`
}

//ReadYml is a method that allows us Read and parse the content of the File
func (x *Config) ReadYml() *Config {
	configFile, err := filepath.Abs("./bot.config.yml")

	if err != nil {
		panic(err)
	}

	yamlFile, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(err)
	}

	ymlError := yaml.Unmarshal(yamlFile, &x)

	if ymlError != nil {
		panic(err)
	}

	return x
}

// GetTokens returns the various tokens declared
func GetTokens() (string, string, string) {
	var c Config

	configObj := c.ReadYml()
	verifyToken, accessToken, appSecret := configObj.VerifyToken, configObj.AccessToken, configObj.AppSecret
	return verifyToken, accessToken, appSecret
}

// GetPersistentMenu returns the persistent menu set by the developer in the config file
func GetPersistentMenu() []messenger.CallToActionsItem {
	var c Config

	configObj := c.ReadYml()
	return configObj.PersistentMenu
}
