package env

import (
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	Port          string `envconfig:"PORT"`
	MysqlUser     string `envconfig:"MYSQL_USER"`
	MysqlPassowrd string `envconfig:"MYSQL_PASSWORD"`
	MysqlAddr     string `envconfig:"MYSQL_ADDR"`
	MysqlDBName   string `envconfig:"MYSQL_DATABASE"`
	AppEnv        string `envconfig:"APP_ENV"`
}

func NewEnvConfig() *EnvConfig {
	var env EnvConfig
	if err := envconfig.Process("", &env); err != nil {
		panic(err.Error())
	}
	return &env
}
