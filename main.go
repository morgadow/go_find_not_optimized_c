package main

import "github.com/morgadow/find_not_optimtimized_c/finder"

var sourceFolder = "C:/workspace/go/src/github.com/morgadow/go_find_not_optimized_c/test"
var sourceFile = "C:/workspace/go/src/github.com/morgadow/go_find_not_optimized_c/test/test_file.c"
var exportFile = "output.txt"

func main() {

	var worker = finder.Tool{}
	//worker.SetSourceFolder(sourceFolder)
	worker.SetSourceFile(sourceFile)
	worker.SetExportFile(exportFile)
	worker.SetAcceptedOptimizations([]string{"Os"})
	worker.FindNonOptimizedFunctions()
}
