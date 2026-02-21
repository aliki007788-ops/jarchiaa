package main

import "os"

type Config struct {
DBHost string
}

func LoadConfig() *Config {
return &Config{
DBHost: getEnv("DB_HOST", "postgres"),
}
}


