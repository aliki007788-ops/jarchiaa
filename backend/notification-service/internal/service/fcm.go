package service

import (
"bytes"
"encoding/json"
"net/http"
)

type FCMService struct {
serverKey string
}

func (s *FCMService) SendPush(token, title, body string) error {
message := map[string]interface{}{
"to": token,
"notification": map[string]string{
"title": title,
"body": body,
},
}

text
jsonData, _ := json.Marshal(message)
req, _ := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewBuffer(jsonData))
req.Header.Set("Authorization", "key="+s.serverKey)

client := &http.Client{}
_, err := client.Do(req)
return err
}


