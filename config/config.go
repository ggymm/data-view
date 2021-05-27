package config

import (
	"os"

	"github.com/pelletier/go-toml"
)

type GlobalConfig struct {
	App     appConfig
	Log     logConfig
	Mysql   mysqlConfig
	Storage storageConfig
}

type appConfig struct {
	Debug            bool
	Addr             string
	AllowCrossDomain bool `toml:"allow_cross_domain"`
}

type logConfig struct {
	Filename string
	MaxSize  int64 `toml:"max_size"`
}

type mysqlConfig struct {
	Address  string
	Port     string
	Name     string
	Username string
	Password string
}

type storageConfig struct {
	Type string
	Url  string
	Path string
}

var Instance *GlobalConfig

func init() {
	var (
		err    error
		config *toml.Tree
	)
	Instance = new(GlobalConfig)
	path := os.Getenv("CONFIG_PATH")
	if len(path) == 0 {
		path = "config.toml"
	}
	config, err = toml.LoadFile(path)
	if err != nil {
		panic(err)
	}
	err = config.Unmarshal(Instance)
	if err != nil {
		panic(err)
	}
}
