package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server Server `yaml:"server"`
	DB     DB     `yaml:"database"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	SslMode  string `yaml:"ssl-mode"`
}

type Server struct {
	Port    string `yaml:"port"`
	Timeout int    `yaml:"timeout"`
}

func (d *DB) GetHost() string {
	return d.Host
}

func (d *DB) GetPort() string {
	return d.Port
}

func (d *DB) GetUser() string {
	return d.User
}

func (d *DB) GetPassword() string {
	return d.Password
}

func (d *DB) GetName() string {
	return d.Name
}

func (d *DB) GetSslMode() string {
	return d.SslMode
}

func New(path string) (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	return &cfg, nil
}
