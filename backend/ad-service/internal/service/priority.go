package service

func CalculatePriority(boost bool, base int) int {
if boost {
return base + 10
}
return base
}


