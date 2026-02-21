package eitaa

// Config تنظیمات کلی ایتا
type Config struct {
	BotToken   string `json:"bot_token"`
	BotUsername string `json:"bot_username"`
	MiniAppURL string `json:"mini_app_url"`
	WebhookURL string `json:"webhook_url"`
}

// DefaultConfig تنظیمات پیش‌فرض
func DefaultConfig() *Config {
	return &Config{
		BotToken:   "your-bot-token",
		BotUsername: "jarchia_bot",
		MiniAppURL: "https://jarchia.ir/eitaa-mini-app",
		WebhookURL: "https://api.jarchia.ir/eitaa/webhook",
	}
}

// GetBotToken دریافت توکن
func (c *Config) GetBotToken() string {
	return c.BotToken
}

// GetMiniAppLink دریافت لینک برنامک
func (c *Config) GetMiniAppLink() string {
	return GenerateMiniAppLink(c.BotUsername)
}
