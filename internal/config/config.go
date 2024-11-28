package config

import (
	"errors"
	"html/template"
	"strconv"
)

type AppConfig struct {
	env                    string
	port                   string
	isTemplateCacheEnabled bool
	templateCache          map[string]*template.Template
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
	_, err := strconv.Atoi(p)

	if err != nil {
		return err
	}

	a.port = p

	return nil
}

func (a *AppConfig) GetPort() string {
	return a.port
}

func (a *AppConfig) EnableTemplateCache() {
	a.isTemplateCacheEnabled = true
}

func (a *AppConfig) DisableTemplateCache() {
	a.isTemplateCacheEnabled = false
}

func (a *AppConfig) IsTemplateCacheEnabled() bool {
	return a.isTemplateCacheEnabled
}

func (a *AppConfig) SetTemplateCache(tc map[string]*template.Template) {
	a.templateCache = tc
}

func (a *AppConfig) GetTemplateCache() map[string]*template.Template {
	return a.templateCache
}
