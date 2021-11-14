package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	O     string = "O"
	O0           = "O0"
	O1           = "O1"
	O2           = "O2"
	O3           = "O3"
	Os           = "Os"
	Ofast        = "Ofast"
	Og           = "Og"
)

var possibleOptimizations = []string{O, O0, O1, O2, O3, Os, Ofast, Og}
var acceptedOptimizations []string

// Function struct representing one function in c code
type Function struct {
	name string
	line int
	file string
}

func (f *Function) ToString() string {
	return fmt.Sprintf("Function | Name: %v, Line: %v, File: %v", f.name, f.line, f.file)
}

// setAcceptedOptimizations set global filter which one is accepted
func setAcceptedOptimizations(optList []string) {
	acceptedOptimizations = optList
}

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

// CleanLine deletes all comments and trailing spaces from c code line
func CleanLine(line string) string {
	// look here: https://yourbasic.org/golang/regexp-cheat-sheet/

	// cut out inline comment strings: /* ... */
	re := regexp.MustCompile("(/\\*.*\\*/)")
	line = re.ReplaceAllString(line, "")

	// cut out multiline comment strings:  * ... (only if not followed by slash which is prohibited by above)
	re = regexp.MustCompile("[ \\t]*\\*.*\\z")
	line = re.ReplaceAllString(line, "")

	// cut out multiline comment strings:  * ... (only if not followed by slash which is prohibited by above)
	re = regexp.MustCompile("[ \\t]*//.*")
	line = re.ReplaceAllString(line, "")

	line = strings.Trim(line, " ")
	line = strings.Trim(line, "\t")
	return line
}

// HasFuncDec checks if line has a function declaration
func HasFuncDec(line string) bool {
	re := regexp.MustCompile("^[ \t]*#")
	if re.MatchString(line) {
		return false
	}
	re = regexp.MustCompile("^[ \\t]*(\\bextern\\b|\\bstatic\\b)?[ \\t]*\\w+[ \\t]+\\w+\\([\\w \\t,]+[,)][{]?$")
	if re.MatchString(line) {
		return true
	}
	return false
}

// HasOptimize checks if line of code contains any optimization: __attribute__((optimize("-Os")))
func HasOptimize(line string) bool {
	line = CleanLine(line) // remove all unwanted shit from line as comments and trailing spaces
	re := regexp.MustCompile("^[ \t]*__attribute__\\(\\(optimize\\(\"-O(.){0,4}\"\\)\\)\\)$")
	if re.MatchString(line) {
		reOpt := regexp.MustCompile("\"-(.){0,4}\"")
		match := reOpt.FindString(line)
		match = strings.Replace(match, "\"-", "", 1)
		match = strings.Replace(match, "\"", "", 1)
		if !elemInSlice(possibleOptimizations, match) {
			fmt.Printf("Invalid optimization found: %s\n", match)
		}
		if elemInSlice(acceptedOptimizations, match) {
			return true
		}
	}
	return false
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

// ExtractFuncName extracts function name from line of valid c code
func ExtractFuncName(line string) string {
	line = strings.Trim(line, " ")
	front := strings.Split(line, "(")[0]
	parts := strings.Split(front, " ")
	return parts[len(parts)-1]
}

// checkFile returns slice with all functions in file which are not optimized yet
func checkFile(lines []string, file string) []Function {
	foundFuncs := make([]Function, 0, 256)
	for idx, elem := range lines {
		line := CleanLine(elem)
		if HasFuncDec(line) {
			if idx > 0 {
				if !HasOptimize(lines[idx-1]) {
					foundFuncs = append(foundFuncs, Function{
						name: ExtractFuncName(line),
						line: idx,
						file: file,
					})
				}
			}
		}
	}
	return foundFuncs
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

// saveToFile saves result to text file
func exportResult(funcs []Function, searchedFiles []string, outputFile, srcFolder, srcFile string) error {

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
	return nil
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

type Worker struct {
	sourceFolder string
	sourceFile   string
	errs         []error
	functions    []Function
	files        []string
}

func (w *Worker) SetSrcFolder(sourceFolder string) {
	if sourceFolder != "" {
		if _, err := os.Stat(sourceFolder); os.IsNotExist(err) {
			fmt.Println("sorce folder does not exist: " + sourceFolder)
			return
		}
		w.sourceFolder = sourceFolder
		w.sourceFile = ""
	}
}


func (w *Worker) SetSrcFile(sourceFile string) {
	if sourceFile != "" {
		if _, err := os.Stat(sourceFile); os.IsNotExist(err) {
			fmt.Println("sorce file does not exist: " + sourceFile)
			return
		}
		w.sourceFolder = ""
		w.sourceFile = sourceFile
	}
}

// FindNonOptimizedFunctions main function to do everything
func (w *Worker) FindNonOptimizedFunctions(srcFolder, srcFile string) []Function {

	// input check and file gathering
	if srcFolder != "" {
		w.files = append(w.files, gatherFiles(srcFolder, ".c", true)...,

		)
	} else if srcFile != "" {
		w.files = append(w.files, srcFile)
	} else {
		w.errs = append(w.errs, errors.New("neither source folder nor source file are given"))
		return nil
	}

	// function gathering
	if len(w.files) == 0 {
		fmt.Println("No file selected for searching")
		return nil
	}
	for _, file := range w.files {
		lines, err := loadFile(file)
		if err != nil {
			w.errs = append(w.errs, err)
		}
		w.functions = append(w.functions, checkFile(lines, file)...)
	}
	err := exportResult(w.functions, w.files, "output.txt", srcFolder, srcFile)
	if err != nil {
		w.errs = append(w.errs, err)
	}
	return w.functions
}
