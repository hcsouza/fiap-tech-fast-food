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
	Port                 string     `mapstructure:"port"`
	AuthConfig           AuthConfig `mapstructure:"authconfig"`
	AuthorizationBaseUrl string     `mapstructure:"authorizationUrl"`
}

type AuthConfig struct {
	UserPoolId string
	ClientId   string
	TokenUse   string
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
		var appConfig Config
		root, _ := find.Repo()
		configFilePath := path.Join(root.Path, "/src/external/api/infra/config")

		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()
		viper.SetConfigName("configs")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(configFilePath)
		viper.AddConfigPath("/app/data/configs")
		err := viper.ReadInConfig()

		if err != nil && !allConfigsAreSet() {
			panic(fmt.Errorf("Falha ao carregar as configurações: %w \n", err))
		}

		if err == nil {
			err := viper.Unmarshal(&appConfig)
			if err != nil {
				panic(err)
			}
		}

		if allConfigsAreSet() { // load envs from infra
			appConfig.ApiCfg.Port = viper.Get("api.port").(string)
			appConfig.MongoCfg.Host = viper.Get("mongodb.host").(string)
			appConfig.MongoCfg.Port = viper.Get("mongodb.port").(string)
			appConfig.MongoCfg.Database = viper.Get("mongodb.database").(string)
		}

		config = &appConfig
	})

	return config
}

func allConfigsAreSet() bool {
	return viper.Get("mongodb.host") != nil &&
		viper.Get("mongodb.port") != nil &&
		viper.Get("mongodb.database") != nil &&
		viper.Get("api.port") != nil
}
