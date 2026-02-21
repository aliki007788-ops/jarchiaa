package service

func CalculateCommission(amount int64, rate float64) int64 {
return int64(float64(amount) * rate / 100)
}


