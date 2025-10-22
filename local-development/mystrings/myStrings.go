package mystrings

// in go Capital name is used to export the function
func Reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}

	return result

}
