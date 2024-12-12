package functionStruct

import "fmt"

type Person struct {
	name string
	age  int
}

func isAbleToVote(person Person) (bool, string) {
	if person.age >= 18 {
		return true, ""
	} else if person.age < 0 {
		return false, "Age is not valid"
	}
	return false, ""
}

func FunctionStruct() {
	person1 := Person{name: "Aditya", age: 18}

	isValid, check := isAbleToVote(person1)

	if len(check) > 0 {
		fmt.Printf("%s", check)
		return
	}

	if isValid {
		fmt.Println("User can vote")
		return
	}
	fmt.Println("User cannot vote")
}
