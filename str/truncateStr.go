package str

// TruncateStr truncate string if length is greater than given length
func TruncateStr(str string, length int) string {
	if len(str) > length {
		return str[0:length] + "..."
	}
	return str
}
