package assignment2

import "fmt"

func isEven(number int) bool {
	//if number%2 == 0 {
	//	return true
	//} else {
	//	return false
	//}

	switch number % 2 {
	case 0:
		return true
	default:
		return false
	}

}

func FindEven() {

	var number int

	fmt.Println("Enter your number:")
	fmt.Scan(&number)

	ans := isEven(number)

	if ans {
		fmt.Println("Your number is even")
		return
	}

	fmt.Println("Your number is odd")
	return

}
