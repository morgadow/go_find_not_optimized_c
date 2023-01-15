package finder

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

const (
	O     string = "O"
	O0    string = "O0"
	O1    string = "O1"
	O2    string = "O2"
	O3    string = "O3"
	Os    string = "Os"
	Ofast string = "Ofast"
	Og    string = "Og"
	Oz    string = "Og"
)

var possibleOptimizations = []string{O, O0, O1, O2, O3, Os, Ofast, Og}

type Tool struct {
	srcDir                string     // file to search, does not search files included here
	srcFile               string     // loads all files in this directory and subdirectories
	exportFile            string     // export file path, if not set a standard file is used
	acceptedOptimizations []string   // optimisations which are accepted to be a valid optimization, this c-function is not returned
	functions             []Function // all found functions with matching oprimization level
	files                 []string   // all files to search
}

func (w *Tool) SetSourceFolder(sourceFolder string) {
	if sourceFolder != "" {
		if _, err := os.Stat(sourceFolder); os.IsNotExist(err) {
			fmt.Println("source folder does not exist: " + sourceFolder)
			return
		}
		w.srcDir = sourceFolder
		w.srcFile = ""
	}
}

func (w *Tool) SetSourceFile(sourceFile string) {
	if sourceFile != "" {
		if _, err := os.Stat(sourceFile); os.IsNotExist(err) {
			fmt.Println("source file does not exist: " + sourceFile)
			return
		}
		w.srcDir = ""
		w.srcFile = sourceFile
	}
}

func (w *Tool) SetExportFile(exportFile string) {
	if exportFile != "" {
		w.exportFile = exportFile
	}
}

// setAcceptedOptimizations set global filter which one is accepted
func (w *Tool) SetAcceptedOptimizations(optList []string) {
	w.acceptedOptimizations = optList
}

// FindNonOptimizedFunctions main function to do everything
func (w *Tool) FindNonOptimizedFunctions() ([]Function, error) {

	// check input
	if w.srcFile == "" && w.srcDir == "" {
		return nil, errors.New("neither source folder nor source file are given")
	}
	if w.exportFile == "" {
		w.exportFile = "optimizeSearchResult.txt"
	}

	// input check and file gathering
	if w.srcDir != "" {
		w.files = append(w.files, gatherFiles(w.srcDir, ".c", true)...)
	}
	if w.srcFile != "" {
		w.files = append(w.files, w.srcFile)
	}

	// check at least one file was found
	if len(w.files) == 0 {
		return nil, errors.New("no file for searching found or selected")
	}

	// search all files
	for _, file := range w.files {
		lines, err := loadFile(file)
		if err != nil {
			return nil, err
		}
		w.functions = append(w.functions, searchFile(lines, file, w.acceptedOptimizations)...)
	}
	err := exportResult(w.functions, w.files, w.acceptedOptimizations, w.exportFile, w.srcDir, w.srcFile)
	if err != nil {
		return nil, err
	}
	return w.functions, nil
}

// saveToFile saves result to text file
func exportResult(funcs []Function, searchedFiles, acceptedOptimizations []string, outputFile, srcFolder, srcFile string) error {

	_, err := os.Stat(outputFile)
	if err == nil {
		err := os.Remove(outputFile)
		if err != nil {
			fmt.Println("Could not delete already existing output file!")
		} else {
			fmt.Println("Deleted already existing output file!")
		}
	}

	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer closeFile(file)
	if err != nil {
		return err
	}
	fmt.Println("Created export file:", file.Name())

	dataWriter := bufio.NewWriter(file)
	defer flushWriter(dataWriter)

	// write header
	currentTime := time.Now()
	_, _ = dataWriter.WriteString("##############################################################\n")
	_, _ = dataWriter.WriteString("Not optimized C Functions  ---  " + currentTime.Format("2006.01.02 15:04:05") + "\n")
	if srcFolder != "" {
		_, _ = dataWriter.WriteString(fmt.Sprintf("Searched root folder: '%s' with %d .c files.\n", srcFolder, len(searchedFiles)))
	}
	if srcFile != "" {
		_, _ = dataWriter.WriteString(fmt.Sprintf("Searched in: '%s'\n", srcFile))
	}
	_, _ = dataWriter.WriteString(fmt.Sprintf("Found %d not optimized functions.\n", len(funcs)))
	_, _ = dataWriter.WriteString(fmt.Sprintf("Accepted optimizations: %s\n", acceptedOptimizations))
	_, _ = dataWriter.WriteString("##############################################################\n\n\n")

	for _, elem := range funcs {
		_, _ = dataWriter.WriteString(elem.ToString() + "\n")
	}
	fmt.Printf("Exported %v functions result to file: %v", len(funcs), outputFile)
	return nil
}
