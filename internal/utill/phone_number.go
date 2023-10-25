package utill

import (
	"regexp"
)

var valid_phone_number_rgx = regexp.MustCompile(`\+998([378]{2}|(9[013-57-9]))\d{7}$`)

func ValidPhoneNumber(pn string) bool {
	if len(pn) < 12 {
		return false
	}
	i := 0
	if pn[0] == '+' {
		i++
	}
	if len(pn)-i != 12 {
		return false
	}
	for ; i < len(pn); i++ {
		if pn[i] < '0' || pn[i] > '9' {
			return false
		}
	}

	return true
}
