package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	v "math-skills/math_skills"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}
	filename := os.Args[1]

	if filename == "" {
		return
	}
	if !strings.HasSuffix(filename, ".txt") {
		fmt.Println("Only .txt files allowed")
		return
	}

	file := v.Readfile(filename)

	Avrg := math.Round(v.Average(file))
	Med := math.Round(v.Median(file))
	Var := math.Round(v.Variance(file))
	StandardDev := math.Round(v.StdDeviation(file))

	fmt.Printf("Average: %.f\n", Avrg)
	fmt.Printf("Median: %.f\n", Med)
	fmt.Printf("Variance: %.f\n", Var)
	fmt.Printf("Standard Deviation: %.f\n", StandardDev)
}
