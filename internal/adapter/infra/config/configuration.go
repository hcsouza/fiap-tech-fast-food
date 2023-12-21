package config

import (
	"fmt"
	"path"
	"strings"
	"sync"

	"github.com/integralist/go-findroot/find"

	"github.com/spf13/viper"
)

var (
	runOnce sync.Once
	config  *Config
)

type Config struct {
	MongoCfg MongoConfig `mapstructure:"mongodb"`
	ApiCfg   Api         `mapstructure:"api"`
}

type MongoConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
}

type Api struct {
	Port string `mapstructure:"port"`
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

		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()
		viper.SetConfigName("configs")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(configFilePath)
		viper.AddConfigPath("/app/data/configs")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Erro fatal no arquivo de configuração: %w \n", err))
		}

		var appConfig Config
		err = viper.Unmarshal(&appConfig)
		if err != nil {
			panic(err)
		}
		config = &appConfig
	})
	return config
}
