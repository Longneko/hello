package config

import (
	"fmt"
	"os"
)

type MySql struct {
	DbName   string
	Password string
}

type Config struct {
	MySql
}

var cfg Config
var cfgInitialized bool

func InitConfig() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("failed to initialize config: %s", err)
		}
	}()

	var isSet bool
	if cfg.MySql.DbName, isSet = os.LookupEnv("HELLO_MYSQL_DATABASE"); !isSet {
		err = fmt.Errorf("env var HELLO_MYSQL_DATABASE is not set")
		return
	}
	if cfg.MySql.Password, isSet = os.LookupEnv("HELLO_MYSQL_ROOT_PASSWORD"); !isSet {
		err = fmt.Errorf("env var HELLO_MYSQL_ROOT_PASSWORD is not set")
		return
	}

	cfgInitialized = true
	return
}

// Get panics if config is not initialized!
func Get() Config {
	if !cfgInitialized {
		panic("config not initialized!")
	}
	return cfg
}
