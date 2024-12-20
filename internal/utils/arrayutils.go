package utils

import "strings"

/*
Check if string array contains string
*/
func Contains(arr []string, str string) bool {
	for _, item := range arr {
		if str == item {
			return true
		}
	}

	return false
}

func ContainsStringInsensitive(slice []string, item string) bool {
	item = strings.ToLower(item)
	for _, str := range slice {
		if strings.ToLower(str) == item {
			return true
		}
	}
	return false
}
