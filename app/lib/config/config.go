package config

import (
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

const (
	// TODO: add a flag for alternate paths. Make platform independent?
	filePath = "./conf/app.conf"

	serverDefaultPort    = "8080"
	serverDefaultReadTO  = 5 * time.Second
	serverDefaultWriteTO = 10 * time.Second

	appDefaultMode = gin.ReleaseMode
)

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

type MySql struct {
	Addr     string
	DbName   string
	User     string
	Password string
}

type Server struct {
	// TODO: add optional host string
	Port         string
	ReadTimeout  duration
	WriteTimeout duration
}

func (s Server) Address() string {
	return fmt.Sprintf(":%s", s.Port)
}

type Application struct {
	Mode string
}

type envOverride func(val string, c *Config) error

var envOverrideMap = map[string]envOverride{
	// TODO: replace with a flag
	"HELLO_APP_MODE": func(envVal string, cfg *Config) error {
		cfg.Application.Mode = envVal
		return nil
	},
	"HELLO_MYSQL_ADDRESS": func(envVal string, cfg *Config) error {
		cfg.MySql.Addr = envVal
		return nil
	},
	"HELLO_MYSQL_DATABASE": func(envVal string, cfg *Config) error {
		cfg.MySql.DbName = envVal
		return nil
	},
	"HELLO_MYSQL_USER": func(envVal string, cfg *Config) error {
		cfg.MySql.User = envVal
		return nil
	},
	"HELLO_MYSQL_PASSWORD": func(envVal string, cfg *Config) error {
		cfg.MySql.Password = envVal
		return nil
	},
	"HELLO_SERVER_PORT": func(envVal string, cfg *Config) error {
		cfg.Server.Port = envVal
		return nil
	},
	"HELLO_SERVER_READ_TO": func(envVal string, cfg *Config) error {
		d, err := time.ParseDuration(envVal)
		if err != nil {
			return err
		}
		cfg.Server.ReadTimeout = duration{d}
		return nil
	},
	"HELLO_SERVER_WRITE_TO": func(envVal string, cfg *Config) error {
		d, err := time.ParseDuration(envVal)
		if err != nil {
			return err
		}
		cfg.Server.WriteTimeout = duration{d}
		return nil
	},
}

type Config struct {
	Application Application
	MySql       MySql
	Server      Server
}

func (c *Config) envOverride(overrides map[string]envOverride) error {
	for envName, f := range overrides {
		envVal, isSet := os.LookupEnv(envName)
		if isSet {
			if err := f(envVal, c); err != nil {
				return fmt.Errorf("error while overriding with env value %s=%s: `%s`",
					envName,
					envVal,
					err,
				)
			}
		}
	}
	return nil
}

var cfg Config
var cfgInitialized bool

func InitConfig() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("failed to initialize config: `%s`", err)
		}
	}()

	// TODO: add filepath parsed from flag
	_, err = toml.DecodeFile(filePath, &cfg)
	if err != nil {
		err = fmt.Errorf("failed to decode config file: `%s`", err)
		return
	}

	err = cfg.envOverride(envOverrideMap)
	if err != nil {
		err = fmt.Errorf("failed to override from environment: `%s`", err)
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
