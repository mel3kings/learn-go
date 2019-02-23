package functions

import "fmt"

type Data struct {
	name string
	age  int
}

func (e Data) changeAge(newAge int) {
	e.age = newAge
}

func (e *Data) changeAge2(newAge int) { // will work
	e.age = newAge
}

func TestPointers() {
	fmt.Println("Testing pointers")
	d := Data{"Mel", 20}
	fmt.Println(d)

	d.changeAge(30) // will not work
	fmt.Println(d)

	d.age = 30 // will work seems pointers are mostly needed on passing through parameters
	fmt.Println(d)

}
