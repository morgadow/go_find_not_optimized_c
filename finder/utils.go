package finder

import (
	"bufio"
	"fmt"
	"os"
)

// elemInSlice helper function which checks if elem is in slice
func elemInSlice(slice []string, key string) bool {
	for _, elem := range slice {
		if elem == key {
			return true
		}
	}
	return false
}

// closeFile helper to defer nicely
func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println(err)
	}
}

// flushWriter helper to defer nicely
func flushWriter(dataWriter *bufio.Writer) {
	err := dataWriter.Flush()
	if err != nil {
		fmt.Println(err)
	}
}

// loadFile returns all lines in file as slice
func loadFile(sourceFile string) ([]string, error) {
	file, err := os.Open(sourceFile)
	defer closeFile(file)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	return txtlines, nil
}

// trunctates string if too long
func trunc(text string, maxLen int) string {
	if len(text) > maxLen {
		return text[:maxLen-5] + " ..."
	}
	return text
}
