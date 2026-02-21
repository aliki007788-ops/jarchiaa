package service

import (
"fmt"
"net/http"
)

type KavenegarService struct {
apiKey string
}

func (s *KavenegarService) SendSMS(phone, message string) error {
url := fmt.Sprintf("https://api.kavenegar.com/v1/%s/sms/send.json", s.apiKey)
resp, err := http.Get(url)
if err != nil {
return err
}
defer resp.Body.Close()
return nil
}


