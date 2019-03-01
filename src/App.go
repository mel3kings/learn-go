package main

import (
	"./data"
	"./functions"
	"fmt"
	"strings"
)

func main() {
	testAllFunctions()
	data.TestMap()
	data.SortEmployeeTest()
}

func testAllFunctions() {
	functions.Imported()
	test, _ := functions.GiveMeData()
	fmt.Println(test)
	functions.RecursiveFibonacci(5)
	fmt.Println(functions.Ok)
	simpleFunction()
	testStringFunctions()
	elegantConstantDeclaration()
	functions.TestPointers()
}

func simpleFunction() {
	//var b rune
	//b = 1234213124
	//fmt.Println(b)
	var x int8 = 5
	var y int8 = 10
	fmt.Println("====== Simple Function: ")
	fmt.Println(x + y)
}

func increment(x int) (int, int) {
	x += 1
	fmt.Println(x)
	return x, 0
}

func basicSyntax() {
	simpleFunction()
	x := 1
	x, y := increment(x)
	fmt.Println(x + y)
}

func testStringFunctions() {
	tester := "Hello World"
	fmt.Println("===== Testing String functions:")
	fmt.Println(tester)
	fmt.Println(len(tester))
	fmt.Println(tester[0:])
	fmt.Println(tester[0:5])

	var hel = `sdfafa asdf aad
		as
		dfafa
		.. as
		as
		dfaf`

	fmt.Println(hel)

	fmt.Println(strings.EqualFold("HELLO", "hello"))
}

func elegantConstantDeclaration() {
	fmt.Println(functions.ErrorCodeTest)
	fmt.Println(functions.ErrorCodeFour)
	fmt.Println(functions.Number2)
}
