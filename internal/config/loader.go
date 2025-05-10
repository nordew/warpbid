package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	instance *Config
	once     sync.Once
)

func Load() (*Config, error) {
	var err error
	once.Do(func() {
		instance = new(Config)
		err = cleanenv.ReadEnv(instance)
	})

	return instance, err
}
