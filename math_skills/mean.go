package math_skills

func Average(numbers []float64) float64 {
	
	sum := 0.00
	for _, number := range numbers {
		sum += number
	}
	return float64(sum) / float64(len(numbers))
}
