package functions

import "fmt"

func RecursiveFibonacci(b int) {
	x, y := 0, 1
	for i := 0; i < b; i++ {
		x, y = y, x+y // tuple declarations
		fmt.Print(x, " ")
	}
}
