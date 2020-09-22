package algos

import (
	"strings"
)

// FindLargestWord finds largest word in the given string
func FindLargestWord(str string) string {
	words := strings.Split(str, " ")
	word := words[0]

	for _, w := range words {
		if len(word) < len(w) {
			word = w
		}
	}

	return word
}
