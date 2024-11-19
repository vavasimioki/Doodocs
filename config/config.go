package config

var ConfigFile = "./config/config.yaml"

type Config struct {
	Port     string `yaml:"port"`     
	Host     string `yaml:"host"`     
	User     string `yaml:"user"`     
	Password string `yaml:"password"` 
}