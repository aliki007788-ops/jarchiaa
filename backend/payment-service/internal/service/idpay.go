package service

import (
"bytes"
"encoding/json"
"net/http"
)

type IDPayService struct {
apiKey string
}

func (s *IDPayService) PaymentRequest(amount int64, callback string) (string, error) {
data := map[string]interface{}{
"order_id": "123",
"amount": amount,
"callback": callback,
}

text
jsonData, _ := json.Marshal(data)
req, _ := http.NewRequest("POST", "https://api.idpay.ir/v1.1/payment", bytes.NewBuffer(jsonData))
req.Header.Set("X-API-KEY", s.apiKey)

client := &http.Client{}
resp, err := client.Do(req)
if err != nil {
    return "", err
}
defer resp.Body.Close()

var result map[string]interface{}
json.NewDecoder(resp.Body).Decode(&result)

return result["id"].(string), nil
}


