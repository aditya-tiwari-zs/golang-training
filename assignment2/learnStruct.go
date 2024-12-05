package assignment2

import "fmt"

type employee struct {
	name string
	age  int
	role string
}

func LearnStruct() {
	var emp1 employee
	emp1.name = "Aditya"
	emp1.age = 24
	emp1.role = "SDE2"

	var poi *employee

	poi = &emp1

	fmt.Printf("%v", *poi)
}
