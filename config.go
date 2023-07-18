package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/configor"
)

type Config struct {
	Server struct {
		Address  string `toml:"address" required:"true"`
		Username string `toml:"username" required:"true"`
		Password string `toml:"password" required:"false"`
		Insecure string `toml:"insecure" required:"false" default:"false"`
	} `toml:"server"`
	Session struct {
		Path string `toml:"path"`
	} `toml:"session"`
}

func NewConfig(path string) (*Config, error) {
	config := &Config{}
	err := configor.Load(config, path)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(config.Session.Path, "~/") {
		config.Session.Path = os.Getenv("HOME") + "/" + strings.TrimPrefix(config.Session.Path, "~/")
	}

	// override configfile supplied password with env variable if one exists
	env_userpass, ok := os.LookupEnv("ZABBIXCTL_USERPASS")
	if ok {
		config.Server.Password = env_userpass
	}
	// must have either a password in the configfile or env variable
	if !ok && config.Server.Password == "" {
		return nil, fmt.Errorf("zabbix user password not found in %s or env variable ZABBIXCTL_USERPASS", path)
	}

	return config, nil
}
