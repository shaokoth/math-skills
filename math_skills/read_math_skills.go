package math_skills

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
// Reads the file input 
func Readfile(filename string) []float64 {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	var result []float64

	hasContent := false
// Scans the file line by line
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		hasContent = true
		line := scanner.Text()
		count++
		values := strings.Split(line, "\n")

		if line == "" {
			continue
		}
		// Converts the string to float64s
		number, err := strconv.ParseFloat(values[0.0], 64)
		if err != nil {
			fmt.Println("Invalid value format in line :",count, err)
			os.Exit(1)
		} else {
			result = append(result, number)
		}
	}
	if len(result) == 0 {
		fmt.Println("The file isempty")
			os.Exit(1)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	if !hasContent {
		fmt.Println("Error: The file is empty")
		os.Exit(1)
	}
	return result
}

