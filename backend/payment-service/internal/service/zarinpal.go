package service

import (
"bytes"
"encoding/json"
"net/http"
)

type ZarinpalService struct {
merchantID string
sandbox bool
}

func (s *ZarinpalService) PaymentRequest(amount int64, callback string) (string, error) {
data := map[string]interface{}{
"MerchantID": s.merchantID,
"Amount": amount,
"CallbackURL": callback,
}

text
jsonData, _ := json.Marshal(data)
resp, err := http.Post("https://api.zarinpal.com/pg/v4/payment/request.json", 
    "application/json", bytes.NewBuffer(jsonData))
if err != nil {
    return "", err
}
defer resp.Body.Close()

var result map[string]interface{}
json.NewDecoder(resp.Body).Decode(&result)

return result["data"].(map[string]interface{})["authority"].(string), nil
}


