package utils

import "strconv"

type Toman int64

func (t Toman) Format() string {
return strconv.FormatInt(int64(t), 10) + " ?????"
}


