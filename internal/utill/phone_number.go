package utill

import "regexp"

var valid_phone_number_rgx = regexp.MustCompile(`\+998([378]{2}|(9[013-57-9]))\d{7}$`)

func ValidPhoneNumber(pn string) bool {
	return valid_phone_number_rgx.MatchString(pn)
}
