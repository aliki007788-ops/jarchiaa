package main

type Config struct {
BotToken string
}

func LoadConfig() *Config {
return &Config{
BotToken: getEnv("EITAA_BOT_TOKEN", ""),
}
}


