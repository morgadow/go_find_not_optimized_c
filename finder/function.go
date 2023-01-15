package finder

import (
	"fmt"
)

// Function struct representing one function in c code
type Function struct {
	name string
	file string
	line int
}

func (f *Function) ToString() string {
	return fmt.Sprintf("Function | Name: %v, Line: %v, File: %v", f.name, f.line, f.file)
}
