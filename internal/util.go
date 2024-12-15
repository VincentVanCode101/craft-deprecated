package internal

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

/*
Get keys of a map as array
*/
func Keys[T any](dict map[string]T) []string {
	keys := make([]string, 0, len(dict))

	for k := range dict {
		keys = append(keys, k)
	}

	return keys
}

/*
Get values of a map as array
*/
func Values(dict map[string]string) []string {
	values := make([]string, 0, len(dict))

	for _, v := range dict {
		values = append(values, v)
	}

	return values
}
