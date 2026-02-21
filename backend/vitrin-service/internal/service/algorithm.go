package service

func CalculateScore(views, likes, saves int64) float64 {
return float64(views)*0.1 + float64(likes)*2 + float64(saves)*1.5
}


