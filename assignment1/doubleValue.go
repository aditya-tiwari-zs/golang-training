package assignment1

import "fmt"

func double(number int) int {
	var ans int
	ans = number * 2
	return ans
}

func DoubleInput() {
	var number int
	fmt.Println("Enter your number which you want to double:")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	result := double(number)
	fmt.Printf("Your numbers's double value is %v \n", result)
}
