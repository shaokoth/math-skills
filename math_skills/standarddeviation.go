package math_skills

import "math"

func StdDeviation(numbers []float64) float64 {
	// Calculates the Standard Deviation
	var squaredsum float64
	mean := 0.00
	n := float64(len(numbers))
	sum := 0.00
	for _, number := range numbers {
		sum += number
	}
	mean = float64(sum) / n

	for _, number := range numbers {
		deviation := float64(number) - mean
		squaredsum += (math.Pow(deviation, 2))
	}
	return math.Sqrt(squaredsum / n)
}
