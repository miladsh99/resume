package utills

import (
	"regexp"
)

func CheckRegexEmail(input string) string {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailPattern)
	if re.MatchString(input) {
		return input
	} else {
		return ""
	}
}
