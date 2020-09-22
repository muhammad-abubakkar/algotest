package algos

import "math"

// CommonMultiple Algo will find common multiple for all integers in the input slice
func CommonMultiple(inputs []int) int {
	min := math.Min(float64(inputs[0]), float64(inputs[1]))
	max := math.Max(float64(inputs[0]), float64(inputs[1]))

	current := max

	for {
		for x := min; x <= max; x++ {
			if math.Mod(current, x) != 0 {
				break
			}

			if x == max {
				return int(current)
			}
		}
		current++
	}
}
