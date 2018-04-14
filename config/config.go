package config

import (
	"sync"

	"github.com/BurntSushi/toml"
)

/**
 * config
 */
type Config struct {
	App      app
	Postgres postgres
}

type app struct {
	Port string
}

type postgres struct {
	Connect string
	MaxIdle int
	MaxOpen int
}

var (
	c    *Config
	once sync.Once
)

/*
 * 返回单例实例
 * @method New
 */
func New() *Config {
	once.Do(func() { //只执行一次
		if _, err := toml.DecodeFile("config.toml", &c); err != nil {
			panic(err.Error())
		}
	})
	return c
}
