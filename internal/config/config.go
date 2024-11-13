package config

import (
	"errors"
	"strconv"
)

type AppConfig struct {
	env string
	port string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

func (a *AppConfig) SetEnv(env string) error {
	if env != "development" && env != "production" {
		return errors.New("invalid env name, valid are : (development or production)")
	}

	a.env = env

	return nil
}

func (a *AppConfig) IsEnvDev() bool {
	return a.env == "development"
}

func (a *AppConfig) SetPort(p string) error {
	_ , err := strconv.Atoi(p)

	if err != nil {
		return err
	}

	a.port = p

	return nil
}

func (a *AppConfig) GetPort() string {
	return a.port
}
