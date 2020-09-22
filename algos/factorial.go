package algos

func fact(input int, result int) int {
	if input <= 0 {
		return result
	}
	return fact(input-1, result*input)
}

// Factorial returns the factorial of the given input
func Factorial(input int) int {
	return fact(input, 1)
}
