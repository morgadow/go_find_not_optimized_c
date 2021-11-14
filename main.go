package main

import (
	uigen "github.com/morgadow/find_not_optimtimized_c/ui"
)



var sourceFolder = "C:\\workspace\\go\\src\\github.com\\morgadow\\find_not_optimtimized_c\\tmp"
var sourceFile = ""

func main() {

	var worker = Worker{
		sourceFolder: sourceFolder,
		sourceFile:   sourceFile,
		errs:         nil,
		functions:    nil,
		files:        nil,
	}
	worker.FindNonOptimizedFunctions(worker.sourceFolder, worker.sourceFile)

	var tool = uigen.UIUiMainWindow{}
	tool.SetupUI()


}
