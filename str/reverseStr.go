package str

// ReverseStr reverse string characters
func ReverseStr(str string) string {
	result := ""
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}

	return result
}
