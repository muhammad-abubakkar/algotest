package str

// RepeatStr repeats given str num times
func RepeatStr(str string, num int) string {
	re := ""
	if num < 1 {
		return re
	}

	for i := 0; i < num; i++ {
		re += str
	}
	return re
}
