package main

type Config struct {
ZarinpalMerchant string
}

func LoadConfig() *Config {
return &Config{
ZarinpalMerchant: getEnv("ZARINPAL_MERCHANT", ""),
}
}


