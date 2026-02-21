package service

import (
"encoding/json"
"net/http"
)

type GeocodingService struct {
apiKey string
}

func (s *GeocodingService) Reverse(lat, lng float64) (string, error) {
resp, err := http.Get("https://api.neshan.org/v5/reverse")
if err != nil {
return "", err
}
defer resp.Body.Close()

text
var result map[string]interface{}
json.NewDecoder(resp.Body).Decode(&result)

return result["formatted_address"].(string), nil
}


