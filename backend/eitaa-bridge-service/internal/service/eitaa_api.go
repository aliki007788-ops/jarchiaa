package service

import (
"bytes"
"encoding/json"
"net/http"
)

type EitaaAPIService struct {
botToken string
}

func (s *EitaaAPIService) SendMessage(chatID, text string) error {
data := map[string]interface{}{
"chat_id": chatID,
"text": text,
}

text
jsonData, _ := json.Marshal(data)
url := "https://eitaayar.ir/api/" + s.botToken + "/sendMessage"

_, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
return err
}


