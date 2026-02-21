package main

type Config struct {
ClickHouseHost string
}

func LoadConfig() *Config {
return &Config{
ClickHouseHost: getEnv("CLICKHOUSE_HOST", "clickhouse"),
}
}


