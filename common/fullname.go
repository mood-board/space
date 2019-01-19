package common

import "strings"

func SplitFullName(fullName string) (string, string) {
	fullName = strings.TrimSpace(fullName)
	firstName := ""
	var lastName string

	i := strings.LastIndex(fullName, " ")

	if i < 0 {
		lastName = fullName
	} else {
		lastName = fullName[i+1:]
		firstName = fullName[:i]
	}

	return strings.TrimSpace(firstName), strings.TrimSpace(lastName)
}
