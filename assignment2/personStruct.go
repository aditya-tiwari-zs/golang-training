package assignment2

import "fmt"

type address struct {
	city    string
	state   string
	pinCode int
}

type person struct {
	name    string
	age     int
	phone   int
	address address
}

func PersonStruct() {

	person1 := person{
		name:  "Aditya",
		age:   24,
		phone: 7007183504,
		address: address{
			city:    "Banglore",
			state:   "Karnataka",
			pinCode: 560102,
		},
	}

	fmt.Println(person1)
	return
}
