package configs

import (
	"encoding/json"
	"os"
)

type Configs struct {
	DataSource         string `json:"data_source"`
	Port               string `json:"port"`
	AddressRedis       string `json:"addressRedis"`
	PasswordRedis      string `json:"passwordRedis"`
	DatabaseredisIndex int    `json:"databaseredisIndex"`
	SecretKey          string `json:"secret_key"`
	ExpireAccess       string `json:"expire_access"`
}

var config *Configs

func Get() *Configs {
	return config
}
func LoadConfig(path string) {
	configFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	byteValue, err := os.ReadFile(configFile.Name())
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		panic(err)
	}
}
