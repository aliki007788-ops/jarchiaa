package service

func CalculatePrice(sizeType string) int64 {
prices := map[string]int64{
"large": 100000,
"medium": 60000,
"small": 30000,
"story": 150000,
}
return prices[sizeType]
}


