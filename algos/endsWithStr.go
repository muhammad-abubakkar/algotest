package algos

// EndsWithStr confirms with str ends with target string
func EndsWithStr(str, target string) bool {
	strLen := len(str)
	tarLen := len(target)
	for idx, ch := range target {
		char := str[strLen-tarLen+idx]
		if rune(char) != ch {
			return false
		}
	}
	return true
}
