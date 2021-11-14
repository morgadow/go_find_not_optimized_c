package main

import (
	"testing"
)


func TestFunction_ToString(t *testing.T) {
	var function = Function{
		name: "testFunction",
		line: 1337,
		file: "just/some/random/test/file.ext",
	}
	var shouldBe = "Function | Name: testFunction, Line: 1337, File: just/some/random/test/file.ext"

	if function.ToString() != shouldBe {
		t.Errorf("Failure: Expected \"%s\", got: \"%s\"", shouldBe, function.ToString())
	}
}

func TestCleanLine(t *testing.T) {

	var linesUnclean = []string{
		"// This is a inline comment",
		" // Same inline comment",
		" \t // Same inline comment with tab",
		"int a = 12; // Random comment",
		"bool func(type param) // Comment",
		"bool func(type param) /* Comment */",
		" bool func(type param) /* Comment */",
		" \tbool func(type param) /* Comment */",
		"// func(type param)",
		" * This is a multiline comment",
		"  * This is a multiline comment with additional steps",
		" \t * This is a multiline comment with tab",
		" int a = 12; /* Comment */ int b = 24;",
	}

	var linesClean = []string{
		"",
		"",
		"",
		"int a = 12;",
		"bool func(type param)",
		"bool func(type param)",
		"bool func(type param)",
		"bool func(type param)",
		"",
		"",
		"",
		"",
		"int a = 12;  int b = 24;",
	}

	if len(linesUnclean) != len(linesClean) {
		t.Errorf("Test and answer array have different lengths! %d, %d", len(linesUnclean), len(linesClean))
	}

	for idx := range linesUnclean {
		if linesClean[idx] != CleanLine(linesUnclean[idx]) {
			t.Errorf("Failure on Index %d!. Expected: %s, got: %s", idx, linesClean[idx], linesUnclean[idx])
		}
	}
}

func TestHasFuncDec(t *testing.T) {

	var linesTest = []string{
		"",
		"bool (func(type param)",
		"bool func(type param",
		"bool func(type param,",
		"bool func(type param)",
		"bool func(type param){",
		"bool func(type param",
		"bool func(void){",
		"bool func(void)",
		"bool func(void,",
		"bool func(void,  ",
		"bool func(void,  \t",
		"// bool func(void,",
		"// bool func(type param,",
		" * bool func(type param,",
		"bool func(type param,",
		"extern bool func(type param,",
		"static bool func(type param,",
		"extern static bool func(type param,",
		"external bool func(type param,",
		"extern \t bool func(type param,",
		"extern \t\t bool funcincation(type param,",
		"extern bool funcincation{type param,",
		"bool funcincation{type param,",
		"bool bool funcincation(type param,",
		"bool funcincation(type param, type param,type param){",
		"\tbool funcincation(type param,",
		"   bool funcincation(type param,",
		"\t bool funcincation(type param,",
		"\t \tbool funcincation(type param,",
	}

	var linesChecked = []bool{
		false,
		false,
		false,
		true,
		true,
		true,
		false,
		true,
		true,
		true,
		false,
		false,
		false,
		false,
		false,
		true,
		true,
		true,
		false,
		false,
		true,
		true,
		false,
		false,
		false,
		true,
		true,
		true,
		true,
		true,
	}

	if len(linesTest) != len(linesChecked) {
		t.Errorf("Test and answer array have different lengths! %d, %d", len(linesTest), len(linesChecked))
	}

	for idx := range linesTest {
		if linesChecked[idx] != HasFuncDec(linesTest[idx]) {
			t.Errorf("Failure on Index %d: %s. Expected: %t, got: %t", idx, linesTest[idx], linesChecked[idx], HasFuncDec(linesTest[idx]))
		}
	}

}

func TestHasOptimize(t *testing.T) {

	setAcceptedOptimizations(possibleOptimizations)

	var linesTest = []string{
		"",
		"// commment __attribute__((optimize(\"-Os\")))",
		"__attribute__((optimize(\"-Os\")))",
		"__attribute__((optimize(\"-Ost\")))",
		"__attribute__((optimize(\"-O5\")))",
		"__attribute__((optimize(\"-0s\")))",
		"__attribute__((optimize(\"-\")))",
		"__attribute__((optimize(\"-O\")))",
		" * __attribute__((optimize(\"-Os\")))",
		"#__attribute__((optimize(\"-Os\")))",
		" \t \t__attribute__((optimize(\"-Os\")))",
		"_attribute__((optimize(\"-Os\")))",
		"__attribute_((optimize(\"-Os\")))",
		"__attribute__(optimize(\"-Os\"))",
		"__attribute__((optimize(\"-Os\"))))",
		"__attribute__((optimize\"-Os\")))",
		"__attribute__ ((optimize-Os\")))",
		"__attribute__ ((optimize-Os)))",
		"__attribute__((optimize-Oss)))\t",
		"__attribute__ ((optimize-Oss)))",
	}

	var linesChecked = []bool{
		false,
		false,
		true,
		false,
		false,
		false,
		false,
		true,
		false,
		false,
		true,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
	}

	if len(linesTest) != len(linesChecked) {
		t.Errorf("Test and answer array have different lengths! %d, %d", len(linesTest), len(linesChecked))
	}

	for idx := range linesTest {
		if linesChecked[idx] != HasOptimize(linesTest[idx]) {
			t.Errorf("Failure on Index %d: %s. Expected: %t, got: %t", idx, linesTest[idx], linesChecked[idx], HasOptimize(linesTest[idx]))
		}
	}
}

func TestExtractFuncName(t *testing.T) {

	var linesTest = []string{
		"bool func(type param",
		"bool func(type param,",
		"bool func(type param)",
		"bool func(type param){",
		"bool func(type param",
		"bool func(void){",
		"bool func(void)",
		"bool func(type param,",
		"extern bool func(type param,",
		"static bool func(type param,",
		"extern \t bool func(type param,",
		"extern \t\t bool funcincation(type param,",
		"bool funcincation(type param, type param,type param){",
		"\tbool funcincation(type param,",
		"   bool funcincation(type param,",
		"\t bool funcincation(type param,",
		"\t \tbool funcincation(type param,",
	}

	var linesChecked = []string{
		"func",
		"func",
		"func",
		"func",
		"func",
		"func",
		"func",
		"func",
		"func",
		"func",
		"func",
		"funcincation",
		"funcincation",
		"funcincation",
		"funcincation",
		"funcincation",
		"funcincation",
	}

	if len(linesTest) != len(linesChecked) {
		t.Errorf("Test and answer array have different lengths! %d, %d", len(linesTest), len(linesChecked))
	}

	for idx := range linesTest {
		if linesChecked[idx] != ExtractFuncName(linesTest[idx]) {
			t.Errorf("Failure on Index %d: %s. Expected: %s, got: %s", idx, linesTest[idx], linesChecked[idx], ExtractFuncName(linesTest[idx]))
		}
	}
}
