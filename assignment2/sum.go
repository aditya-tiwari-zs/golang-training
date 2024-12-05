package assignment2

import "fmt"

func SumAll() {
	var number int
	res := 0
	fmt.Println("Enter a number:")
	fmt.Scan(&number)
	for i := 1; i <= number; i++ {
		res = res + i
	}

	fmt.Printf("The sum of number from 1 to %d is %d", number, res)
	return
}
