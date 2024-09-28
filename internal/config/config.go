package config

import "time"

type Config struct {
	Env            string `yaml:"env" env-default:"development"`
	StoragePath    string `yaml:"storage_path" env-required:"true"`
	HTTPServer     `yaml:"http_server"`
	DataBaseConfig `yaml:"db"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"0.0.0.0:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type DataBaseConfig struct {
	Address  string `yaml:"address" env-default:"0.0.0.0:2000"`
	User     string `yaml:"user" env-default:"admin"`
	Password string `yaml:"password" env-default:"admin"`
}
