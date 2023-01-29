package finder

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
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

var PossibleOptimizations = []string{O, O0, O1, O2, O3, Os, Ofast, Og}

const MaxStatusLen = 60 // status strings will be truncated if longer than this
var folderCnt = 0       // status indicator for how many folders are searched

type Tool struct {
	srcDir                string     // file to search, does not search files included here
	srcFile               string     // loads all files in this directory and subdirectories
	exportFile            string     // export file path, if not set a standard file is used
	status                string     // current status of search, indicates which file is searched
	progress              float64    // present progress of search
	acceptedOptimizations []string   // optimisations which are accepted to be a valid optimization, this c-function is not returned
	functions             []Function // all found functions with matching oprimization level
	files                 []string   // all files to search
}

func NewTool() *Tool {
	return &Tool{}
}

func (w *Tool) GetStatus() string {
	return w.status
}

func (w *Tool) GetProgress() float64 {
	return w.progress
}

func (w *Tool) SetSourceFolder(sourceFolder string) {
	if sourceFolder != "" {
		if _, err := os.Stat(sourceFolder); os.IsNotExist(err) {
			fmt.Println("source folder does not exist: " + sourceFolder)
			return
		}
		w.srcDir = sourceFolder
		w.srcFile = ""
		fmt.Println("set source folder to: ", sourceFolder)
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
		fmt.Println("set source file to: ", sourceFile)
	}
}

func (w *Tool) GetSourceFile() string {
	return w.srcFile
}

func (w *Tool) GetSourceFolder() string {
	return w.srcDir
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
	w.progress = 0
	w.status = ""
	folderCnt = 0
	if w.srcFile == "" && w.srcDir == "" {
		return nil, errors.New("neither source folder nor source file are given")
	}
	if w.exportFile == "" {
		dir, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		w.exportFile = filepath.Join(dir, "optimizeSearchResult.txt")
	}

	// input check and file gathering
	if w.srcDir != "" {
		w.status = fmt.Sprintf("Gathering from %v: %v", folderCnt, trunc(w.GetSourceFolder(), MaxStatusLen-20))
		w.files = append(w.files, w.gatherFiles(w.srcDir, ".c", true)...)
	}
	if w.srcFile != "" {
		w.files = append(w.files, w.srcFile)
	}
	w.progress = 0.05

	// check at least one file was found
	if len(w.files) == 0 {
		w.status = ""
		w.progress = 0
		return nil, errors.New("no file for searching found in selected files and folders")
	}

	// search all files
	var incr float64 = 0.9 / float64(len(w.files)) // should end at 95 %
	for idx, file := range w.files {
		w.status = fmt.Sprintf("Analyzing file %v: %v", idx, trunc(file, MaxStatusLen-20))
		lines, err := loadFile(file)
		if err != nil {
			return nil, err
		}
		w.functions = append(w.functions, searchFile(lines, file, w.acceptedOptimizations)...)
		w.progress += float64(incr)
	}
	w.status = fmt.Sprintf("Exporting result: %v", trunc(w.exportFile, MaxStatusLen-22))
	err := exportResult(w.functions, w.files, w.acceptedOptimizations, w.exportFile, w.srcDir, w.srcFile)
	if err != nil {
		return nil, err
	}
	w.status = "Done ..."
	w.progress = 1
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

// gatherFiles recursive function to search for files with matching file extension
func (w *Tool) gatherFiles(folder, ext string, inclSubs bool) []string {
	folderCnt += 1 // increase folder count by one
	files := make([]string, 0, 256)
	elems, _ := os.ReadDir(folder)
	for _, elem := range elems {
		if elem.IsDir() && inclSubs {
			w.status = fmt.Sprintf("Gathering files %v: %v", folderCnt, trunc(filepath.Join(folder, elem.Name()), MaxStatusLen-23))
			files = append(files, w.gatherFiles(path.Join(folder, elem.Name()), ext, inclSubs)...)
		} else if ext != "" && filepath.Ext(elem.Name()) == ext {
			files = append(files, path.Join(folder, elem.Name()))
		}
	}
	return files
}
