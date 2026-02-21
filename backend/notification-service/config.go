package main

type Config struct {
FCMKey string
}

func LoadConfig() *Config {
return &Config{
FCMKey: getEnv("FCM_SERVER_KEY", ""),
}
}


