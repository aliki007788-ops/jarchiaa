package main

type Config struct {
DBHost string
DBPort string
DBName string
DBUser string
DBPass string
}

func LoadConfig() *Config {
return &Config{
DBHost: getEnv("DB_HOST", "postgres"),
DBPort: getEnv("DB_PORT", "5432"),
DBName: getEnv("DB_NAME", "jarchia"),
DBUser: getEnv("DB_USER", "jarchia"),
DBPass: getEnv("DB_PASSWORD", ""),
}
}


