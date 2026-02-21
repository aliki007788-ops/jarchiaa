package utils

import "time"

type JalaliDate struct {
Year int
Month int
Day int
}

func NowJalali() JalaliDate {
t := time.Now()
return GregorianToJalali(t)
}


