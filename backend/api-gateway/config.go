package main

type Config struct {
Port string
JWTSecret string
}

func LoadConfig() *Config {
return &Config{
Port: getEnv("PORT", "8080"),
JWTSecret: getEnv("JWT_SECRET", "default-secret"),
}
}

func getEnv(key, defaultValue string) string {
if value := os.Getenv(key); value != "" {
return value
}
return defaultValue
}


