package functions

import "fmt"

func Imported() {
	fmt.Println("hello from imported")
}

func GiveMeData() (string, int) {
	x := 1
	return "test", x
}
