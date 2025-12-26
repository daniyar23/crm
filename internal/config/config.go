package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	User     string `env:"DB_USER" required:"true"`
	Password string `env:"DB_PASSWORD" required:"true"`
	Host     string `env:"DB_HOST" required:"true"`
	Port     int    `env:"DB_PORT" required:"true"`
	Name     string `env:"DB_NAME" required:"true"`
}

func (c DBConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
	)
}

func Load() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
