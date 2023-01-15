package finder

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// gatherFiles recursive function to search for files with matching file extension
func gatherFiles(folder, ext string, inclSubs bool) []string {
	files := make([]string, 0, 256)
	elems, _ := os.ReadDir(folder)
	for _, elem := range elems {
		if elem.IsDir() && inclSubs {
			files = append(files, gatherFiles(path.Join(folder, elem.Name()), ext, inclSubs)...)
		} else if ext != "" && filepath.Ext(elem.Name()) == ext {
			files = append(files, path.Join(folder, elem.Name()))
		}
	}
	return files
}

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
