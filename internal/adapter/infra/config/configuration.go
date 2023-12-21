package config

import (
	"encoding/json"
	"fmt"
	"path"
	"sync"

	"github.com/integralist/go-findroot/find"

	"github.com/spf13/viper"
)

var (
	runOnce sync.Once
	config  *Config
)

type Config struct {
	MongoCfg MongoConfig `json:"mongodb"`
	ApiCfg   Api         `json:"api"`
}

type MongoConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Api struct {
	Port string `json:"port"`
}

func init() {
	config = setupConfig()
}

func GetMongoCfg() MongoConfig {
	return config.MongoCfg
}

func GetApiCfg() Api {
	return config.ApiCfg
}

func setupConfig() *Config {

	runOnce.Do(func() {
		root, _ := find.Repo()
		configFilePath := path.Join(root.Path, "/internal/adapter/infra/config")

		viper.AutomaticEnv()
		viper.SetConfigName("configs")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(configFilePath)
		viper.AddConfigPath("/app/data/configs")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Erro fatal no arquivo de configuração: %w \n", err))
		}

		b, err := json.Marshal(viper.Get("configuration"))
		if err != nil {
			panic(err)
		}

		var appConfig Config
		err = json.Unmarshal(b, &appConfig)
		if err != nil {
			panic(err)
		}
		config = &appConfig
	})
	return config
}
