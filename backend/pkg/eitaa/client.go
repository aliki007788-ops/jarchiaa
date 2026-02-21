package eitaa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// EitaaClient کلاینت اصلی برای ارتباط با ایتا
type EitaaClient struct {
	BotToken string
	BaseURL  string
	client   *http.Client
}

// NewEitaaClient ایجاد کلاینت جدید
func NewEitaaClient(token string) *EitaaClient {
	return &EitaaClient{
		BotToken: token,
		BaseURL:  "https://eitaayar.ir/api",
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// MessageRequest ساختار پیام
type MessageRequest struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

// SendMessage ارسال پیام ساده
func (c *EitaaClient) SendMessage(chatID, text string) error {
	url := fmt.Sprintf("%s/%s/sendMessage", c.BaseURL, c.BotToken)
	
	reqBody := MessageRequest{
		ChatID: chatID,
		Text:   text,
	}
	
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("خطا در ساخت JSON: %v", err)
	}
	
	resp, err := c.client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("خطا در ارسال درخواست: %v", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("خطا از سمت ایتا: %s", resp.Status)
	}
	
	return nil
}

// SendAdShare ارسال آگهی به ایتا
func (c *EitaaClient) SendAdShare(chatID, adTitle, adDesc, adLink string) error {
	message := fmt.Sprintf("📢 *%s*\n\n%s\n\n🔗 %s", adTitle, adDesc, adLink)
	return c.SendMessage(chatID, message)
}
