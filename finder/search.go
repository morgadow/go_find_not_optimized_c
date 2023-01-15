package finder

import (
	"fmt"
	"regexp"
	"strings"
)

// look here: https://yourbasic.org/golang/regexp-cheat-sheet/

// searchFile returns slice with all functions in file which are not optimized yet
func searchFile(lines []string, file string, acceptedOptimizations []string) []Function {
	foundFuncs := make([]Function, 0, 256)
	for idx, elem := range lines {
		line := CleanLine(elem)
		if HasFuncDec(line) {
			if idx > 0 {
				if !HasOptimize(lines[idx-1], acceptedOptimizations) {
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

// HasFuncDec checks if line has a function declaration, eg.: extern bool funcname(uint8 value){
func HasFuncDec(line string) bool {
	var re = regexp.MustCompile(`^(\bextern\b|\bstatic\b)?[ \t]*\w+[ \t]*\w+[ \t]*\(((\w*[ \t]*\w*[,]?)*|[(void) \t])[,)][{]?$`)
	return re.MatchString(line)
}

// HasOptimize checks if line of code contains any optimization: __attribute__((optimize("-Os")))
func HasOptimize(line string, acceptedOptimizations []string) bool {
	line = CleanLine(line) // remove all unwanted shit from line as comments and trailing spaces
	var reOpt = regexp.MustCompile(`^__attribute__\(\(optimize\(\"-O(.){0,4}\"\)\)\)$`)
	if reOpt.MatchString(line) {

		// extract optimization after checking line has any
		reOpt := regexp.MustCompile(`\"-(.){0,4}\"`)
		match := reOpt.FindString(line)
		match = strings.Replace(match, "\"-", "", 1)
		match = strings.Replace(match, "\"", "", 1)

		// check optimization is in ignore list
		if !elemInSlice(possibleOptimizations, match) {
			fmt.Printf("Invalid optimization found: %s\n", match)
		}
		if elemInSlice(acceptedOptimizations, match) {
			return true
		}
	}
	return false
}

// CleanLine deletes all comments and trailing spaces from c code line
func CleanLine(line string) string {

	// cut out all inline comment strings: // ...
	re := regexp.MustCompile(`//.*`)
	line = re.ReplaceAllString(line, "")

	// cut out inline comment strings: /* ... */
	re = regexp.MustCompile(`(/\\*.*\\*/)`)
	line = re.ReplaceAllString(line, "")

	// cut out startline of multiline comment strings:  \* ... (only if not followed by slash which is prohibited by above)
	re = regexp.MustCompile(`(/\*.*)`)
	line = re.ReplaceAllString(line, "")

	// cut out multiline comment strings:  * ... (only if not followed by slash which is prohibited by above)
	re = regexp.MustCompile(`(\*.*)`)
	line = re.ReplaceAllString(line, "")

	// clean out whitespaces and tabs at end and start of string
	re = regexp.MustCompile(`(^[ \t]*)`)
	line = re.ReplaceAllString(line, "")
	re = regexp.MustCompile(`([ \t]*$)`)
	line = re.ReplaceAllString(line, "")

	return line
}

// ExtractFuncName extracts function name from line of valid c code
func ExtractFuncName(line string) string {
	line = strings.Trim(line, " ")
	front := strings.Split(line, "(")[0]
	parts := strings.Split(front, " ")
	return parts[len(parts)-1]
}
