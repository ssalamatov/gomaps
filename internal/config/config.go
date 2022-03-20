package config

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

const FileConfigName string = "config.yaml"

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Postgresql struct {
		Host string `yaml:"host"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		DB   string `yaml:"db"`
		Port string `yaml:"port"`
	} `yaml:"postgresql"`
}

func NewConfig() (*Config, error) {
	root, ok := os.LookupEnv("GOMAPSROOT")
	if !ok {
		return nil, NewErrMissingEnv(nil)
	}

	configPath := path.Join(root, "config", FileConfigName)
	return ReadConfig(configPath)
}

func (config *Config) GetAppAddr() string {
	host := config.Server.Host
	port := config.Server.Port
	return fmt.Sprintf("%s:%s", host, port)
}
func (config *Config) GetDbDsn() string {
	user := config.Postgresql.User
	pass := config.Postgresql.Pass
	db := config.Postgresql.DB
	host := config.Postgresql.Host
	port := config.Postgresql.Port
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, db)
}

func ReadConfig(configPath string) (*Config, error) {
	config := &Config{}

	f, err := os.Open(configPath)
	if err != nil {
		return nil, NewErrConfigParseFailed(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&config); err != nil {
		return nil, NewErrConfigNotFound(err)
	}
	return config, nil
}
