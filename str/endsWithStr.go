package str

// EndsWithStr confirms with str ends with target string
func EndsWithStr(str, target string) bool {
	strLen := len(str)
	tarLen := len(target)
	chunk := str[strLen-tarLen:]
	return chunk == target
}
