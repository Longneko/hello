package config

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	serverDefaultPort    = "8080"
	serverDefaultReadTO  = 5 * time.Second
	serverDefaultWriteTO = 10 * time.Second

	appDefaultMode = gin.ReleaseMode
)

type MySql struct {
	DbName   string
	Password string
}

type Server struct {
	// TODO: add optional host string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (s Server) Address() string {
	return fmt.Sprintf(":%s", s.Port)
}

type Application struct {
	Mode string
}

type Config struct {
	Application
	MySql
	Server
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
	
	// TODO: replace with a flag
	if cfg.Application.Mode, isSet = os.LookupEnv("HELLO_APP_MODE"); !isSet {
		cfg.Application.Mode = appDefaultMode
		fmt.Printf("config: application mode not defined, using default `%s`\n", cfg.Application.Mode)
	}

	// MySQL
	if cfg.MySql.DbName, isSet = os.LookupEnv("HELLO_MYSQL_DATABASE"); !isSet {
		err = fmt.Errorf("env var HELLO_MYSQL_DATABASE is not set")
		return
	}
	if cfg.MySql.Password, isSet = os.LookupEnv("HELLO_MYSQL_ROOT_PASSWORD"); !isSet {
		err = fmt.Errorf("env var HELLO_MYSQL_ROOT_PASSWORD is not set")
		return
	}

	// Gin-Gonic server
	cfg.Server.Port, isSet = os.LookupEnv("HELLO_SERVER_PORT")
	if !isSet {
		cfg.Server.Port = serverDefaultPort
		fmt.Printf("config: server port not defined, using default `%s`\n", cfg.Server.Port)
	}
	readTOStr, isSet := os.LookupEnv("HELLO_SERVER_READ_TO")
	if isSet {
		cfg.Server.ReadTimeout, err = time.ParseDuration(readTOStr)
		if err != nil {
			err = fmt.Errorf("failed to partse duration from 'HELLO_SERVER_READ_TO' var. Orig. Err: `%s`", err)
			return
		}
	} else {
		cfg.Server.ReadTimeout = serverDefaultReadTO
		fmt.Printf("config: server read timeout not defined, using default `%s`\n", cfg.Server.ReadTimeout)
	}
	writeTOStr, isSet := os.LookupEnv("HELLO_SERVER_WRITE_TO")
	if isSet {
		cfg.Server.WriteTimeout, err = time.ParseDuration(writeTOStr)
		if err != nil {
			err = fmt.Errorf("failed to partse duration from 'HELLO_SERVER_WRITE_TO' var. Orig. Err: `%s`", err)
			return
		}
	} else {
		cfg.Server.WriteTimeout = serverDefaultWriteTO
		fmt.Printf("config: server write timeout not defined, using default `%s`\n", cfg.Server.WriteTimeout)
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
