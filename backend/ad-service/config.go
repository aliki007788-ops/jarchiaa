package main

type Config struct {
Port string
}

func LoadConfig() *Config {
return &Config{
Port: getEnv("PORT", "8083"),
}
}


