package service

import "strings"

func ModifyValue(v string) string {
	newvaluse := strings.TrimSpace(strings.ToLower(v))
	return newvaluse
}

func ModifyPassword(p []byte) string {
	newvaluse := strings.ToLower(string(p))
	return newvaluse
}
