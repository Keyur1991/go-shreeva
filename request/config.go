package request

import (
	"github.com/Keyur1991/go-shreeva/config"
)

type Getters interface {
	GetRateLimit()
	GetCors()
}

type Config struct {
	Cors *config.Cors
	RateLimit *config.RateLimit
}

func (conf *Config) GetRateLimit() *config.RateLimit {
	return conf.RateLimit
}

func (conf *Config) GetCors() *config.Cors {
	return conf.Cors
}