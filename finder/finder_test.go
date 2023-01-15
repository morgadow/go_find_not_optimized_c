package finder

import (
	"testing"
)

func TestFunction_elemInSlice(t *testing.T) {
	var elem string = "elem"
	var slice = []string{"elem", "foo", "bar"}

	if !elemInSlice(slice, elem) {
		t.Errorf("Failure: Expected to find element 'elem' in slice!")
	}

	elem = "elem"
	slice = []string{"foo", "bar"}

	if elemInSlice(slice, elem) {
		t.Errorf("Failure: Expected to not find element 'elem' in slice!")
	}
}

func TestFunction_ToString(t *testing.T) {
	var function = Function{
		name: "testFunction",
		line: 1337,
		file: "just/some/random/test/file.ext",
	}
	var shouldBe = "Function | Name: testFunction, Line: 1337, File: just/some/random/test/file.ext"

	if function.ToString() != shouldBe {
		t.Errorf("Failure: Expected \"%s\", got: : \"%s\"", shouldBe, function.ToString())
	}
}

func TestCleanLine(t *testing.T) {

	var idx = 1
	auxTest := func(line, exp string) {
		ret := CleanLine(line)
		if ret != exp {
			t.Errorf("%v: Expected line '%s' to have answer: '%v', got: '%v'", idx, line, exp, ret)
		}
		idx++
	}

	auxTest("// This is a inline comment", "")
	auxTest(" // Same inline comment", "")
	auxTest(" \t // Same inline comment with tab", "")
	auxTest("int a = 12; // Random comment", "int a = 12;")
	auxTest("bool func(type param) // Comment", "bool func(type param)")
	auxTest("bool func(type param) /* Comment */", "bool func(type param)")
	auxTest(" bool func(type param) /* Comment */", "bool func(type param)")
	auxTest(" \tbool func(type param) /* Comment */", "bool func(type param)")
	auxTest("// func(type param)", "")
	auxTest(" * This is a multiline comment", "")
	auxTest(" /* This is the startline of a multiline comment", "")
	auxTest("  * This is a multiline comment with additional steps", "")
	auxTest(" \t * This is a multiline comment with tab", "")
	auxTest(" int a = 12; /* Comment */ int b = 24;", "int a = 12;  int b = 24;")
}

func TestHasFuncDec(t *testing.T) {

	var idx = 1
	auxTest := func(line string, exp bool) {
		ret := HasFuncDec(line)
		if ret != exp {
			t.Errorf("%v: Expected line '%s' to have answer: %v, got: %v", idx, line, exp, ret)
		}
		idx++
	}

	auxTest("", false)
	auxTest("bool (func(type param)", false)
	auxTest("bool func(type param", false)
	auxTest("bool func(type param,", true)
	auxTest("bool func(type param)", true)
	auxTest("bool func(type param){", true)
	auxTest("bool func(type param", false)
	auxTest("bool func(void){", true)
	auxTest("bool func(void)", true)
	auxTest("bool func(uint8 param_1, void)", true) // not really okay, but to tidious if i wanted to fix that
	auxTest("bool func(void,  ", false)
	auxTest("bool func(void,  \t", false)
	auxTest("bool func(uint8 param_1, void)", true) // not really okay, but to tidious if i wanted to fix that
	auxTest("bool func(uint8 param_1)", true)
	auxTest("bool func()", true)
	auxTest("bool func( )", true)
	auxTest("bool func(void)", true)
	auxTest("bool func(void,", true) // not really okay, but to tidious if i wanted to fix that
	auxTest("bool func(uint8 param_1,", true)
	auxTest("// bool func(void,", false)
	auxTest("// bool func(type param,", false)
	auxTest(" * bool func(type param,", false)
	auxTest("bool func(type param,", true)
	auxTest("extern bool func(type param,", true)
	auxTest("static bool func(type param,", true)
	auxTest("extern static bool func(type param,", false)
	auxTest("external bool func(type param,", false)
	auxTest("static \t bool func(type param,", true)
	auxTest("extern \t bool func(type param,", true)
	auxTest("extern \t\t bool funcincation(type param,", true)
	auxTest("extern bool funcincation{type param,", false)
	auxTest("bool funcincation{type param,", false)
	auxTest("bool bool funcincation(type param,", false)
	auxTest("bool funcincation(type param, type param,type param){", true)
	auxTest("\tbool funcincation(type param,", true)
	auxTest("   bool funcincation(type param,", true)
	auxTest("\t bool funcincation(type param,", true)
	auxTest("\t \tbool funcincation(type param,", true)

}

func TestHasOptimize(t *testing.T) {

	var idx = 1
	testFunc := func(line string, exp bool) {
		ret := HasOptimize(line, possibleOptimizations)
		if ret != exp {
			t.Errorf("%v: Expected line '%s' to have answer: %v, got: %v", idx, line, exp, ret)
		}
		idx++
	}

	testFunc("", false)
	testFunc("// commment __attribute__((optimize(\"-Os\")))", false)
	testFunc("__attribute__((optimize(\"-Os\")))", true)
	testFunc("__attribute__((optimize(\"-Ost\")))", false)
	testFunc("__attribute__((optimize(\"-O5\")))", false)
	testFunc("__attribute__((optimize(\"-0s\")))", false)
	testFunc("__attribute__((optimize(\"-\")))", false)
	testFunc("__attribute__((optimize(\"-O\")))", true)
	testFunc(" * __attribute__((optimize(\"-Os\")))", false)
	testFunc("#__attribute__((optimize(\"-Os\")))", false)
	testFunc(" \t \t__attribute__((optimize(\"-Os\")))", true)
	testFunc("_attribute__((optimize(\"-Os\")))", false)
	testFunc("__attribute_((optimize(\"-Os\")))", false)
	testFunc("__attribute__(optimize(\"-Os\"))", false)
	testFunc("__attribute__((optimize(\"-Os\"))))", false)
	testFunc("__attribute__((optimize\"-Os\")))", false)
	testFunc("__attribute__ ((optimize-Os\")))", false)
	testFunc("__attribute__ ((optimize-Os)))", false)
	testFunc("__attribute__((optimize-Oss)))\t", false)
	testFunc("__attribute__ ((optimize-Oss)))", false)

}

func TestExtractFuncName(t *testing.T) {

	testFunc := func(line string, exp string) {
		ret := ExtractFuncName(line)
		if ret != exp {
			t.Errorf("Expected line '%s' to have answer: %v, got: %v", line, exp, ret)
		}
	}

	testFunc("bool func(type param", "func")
	testFunc("bool func(type param,", "func")
	testFunc("bool func(type param)", "func")
	testFunc("bool func(type param){", "func")
	testFunc("bool func(type param", "func")
	testFunc("bool func(void){", "func")
	testFunc("bool func(void)", "func")
	testFunc("bool func(type param,", "func")
	testFunc("extern bool func(type param,", "func")
	testFunc("static bool func(type param,", "func")
	testFunc("extern \t bool func(type param,", "func")
	testFunc("extern \t\t bool funcincation(type param,", "funcincation")
	testFunc("bool funcincation(type param, type param,type param){", "funcincation")
	testFunc("\tbool funcincation(type param,", "funcincation")
	testFunc("   bool funcincation(type param,", "funcincation")
	testFunc("\t bool funcincation(type param,", "funcincation")
	testFunc("\t \tbool funcincation(type param,", "funcincation")
}
