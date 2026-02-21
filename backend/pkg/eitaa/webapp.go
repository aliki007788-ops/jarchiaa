package eitaa

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WebAppData داده‌های دریافتی از برنامک ایتا
type WebAppData struct {
	QueryID  string `json:"query_id"`
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	AuthDate int64  `json:"auth_date"`
	Hash     string `json:"hash"`
}

// MiniAppConfig تنظیمات برنامک ایتا
type MiniAppConfig struct {
	AppName     string `json:"app_name"`
	ButtonText  string `json:"button_text"`
	StartMessage string `json:"start_message"`
}

// GenerateMiniAppLink ساخت لینک برنامک
func GenerateMiniAppLink(username string) string {
	return fmt.Sprintf("https://eitaa.com/%s", username)
}

// ValidateWebApp بررسی صحت درخواست از برنامک
func ValidateWebApp(r *http.Request) bool {
	// بررسی هدر مخصوص ایتا
	userAgent := r.Header.Get("User-Agent")
	return userAgent != "" && (userAgent == "EitaaBot" || userAgent == "Eitaa")
}

// ParseWebAppData解析 داده‌های دریافتی از برنامک
func ParseWebAppData(data string) (*WebAppData, error) {
	var webAppData WebAppData
	err := json.Unmarshal([]byte(data), &webAppData)
	if err != nil {
		return nil, err
	}
	return &webAppData, nil
}
