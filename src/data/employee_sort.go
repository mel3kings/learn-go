package data

import (
	"fmt"
	"sort"
)

type Employee struct {
	employeeId   int
	employeeName string
}

type byEmployeeId []Employee

func (a byEmployeeId) Len() int           { return len(a) }
func (a byEmployeeId) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byEmployeeId) Less(i, j int) bool { return a[i].employeeId < a[j].employeeId }

func SortEmployeeTest() {
	employee10 := Employee{10, "mel"}
	employee2 := Employee{2, "april"}
	employee3 := Employee{6, "dummy"}
	listOfEmployees := []Employee{employee10, employee3, employee2}
	fmt.Println(listOfEmployees)
	sort.Sort(byEmployeeId(listOfEmployees))
	fmt.Println(listOfEmployees)
}
