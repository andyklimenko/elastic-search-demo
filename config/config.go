package config

import (
	"errors"
	"strings"

	"github.com/spf13/viper"
)

var (
	ErrNoElasticPassword = errors.New("no elastic password")
	ErrNoElasticAddress  = errors.New("no elastic address")
)

type Config struct {
	Elastic Elastic
}

type Elastic struct {
	UserName string
	Password string
	Address  string
}

func (e *Elastic) load(envPrefix string) error {
	v := setupViper(envPrefix)

	v.SetDefault("username", "elastic")
	e.UserName = v.GetString("username")
	e.Password = v.GetString("password")
	if e.Password == "" {
		return ErrNoElasticPassword
	}

	e.Address = v.GetString("address")
	if e.Address == "" {
		return ErrNoElasticAddress
	}

	return nil
}

func setupViper(envPrefix string) *viper.Viper {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetEnvPrefix(envPrefix)
	return v
}

func Load() (Config, error) {
	var e Elastic
	if err := e.load("elastic"); err != nil {
		return Config{}, err
	}

	return Config{Elastic: e}, nil
}
