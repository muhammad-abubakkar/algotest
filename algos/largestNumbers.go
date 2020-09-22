package algos

// FindLargestNumbers finds largest number in array of numbers
func FindLargestNumbers(nums [][]int) []int {
	var largest []int

	for _, group := range nums {
		n := group[0]
		for _, num := range group {
			if n < num {
				n = num
			}
		}
		largest = append(largest, n)
	}

	return largest
}
