package utils

import "regexp"

func ValidatePhone(phone string) bool {
pattern := ^09[0-9]{9}$
matched, _ := regexp.MatchString(pattern, phone)
return matched
}


