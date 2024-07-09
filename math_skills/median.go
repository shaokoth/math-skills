package math_skills

import "sort"

func Median(numbers []float64) float64 {
	// Finds the middle value after rearranging in an ascending or descending order
	sort.Float64s(numbers)
	n := len(numbers)
	if n%2 == 1 {
		return float64(numbers[n/2])
	}
	mid1 := numbers[n/2]
	mid2 := numbers[(n+1)/2]
	return float64(mid1+mid2) / 2
}
