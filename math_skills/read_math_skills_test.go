package math_skills

import (
    "os"
    "reflect"
    "testing"
)

func createTestFile(filename string, content string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    _, err = file.WriteString(content)
    return err
}

func TestReadfile(t *testing.T) {
    // Define the test data content
    testData := `1.0, 2.5, 3.7, f, 4.2`

    // Create a temporary file with the test data
    testFilename := "test_data.txt"
    err := createTestFile(testFilename, testData)
    if err != nil {
        t.Fatalf("Failed to create test file: %v", err)
    }
    defer os.Remove(testFilename)

    // Call the Readfile function with the test file path
    result := Readfile(testFilename)

    // Define the expected result based on the valid lines in the test file
    expected := []float64{1.0, 2.5, 3.7, 4.2}

    // Compare the result with the expected value
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Readfile(%s) = %v, want %v", testFilename, result, expected)
    }
}
