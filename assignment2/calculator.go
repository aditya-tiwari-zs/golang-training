package assignment2

import (
	"fmt"
)

func calculate(x, y float32, operator string) float32 {
	switch operator {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		{
			if y == 0 {
				panic("division with 0 is not supported")
			}
			return x / y
		}
	default:
		panic("this operator is not supported")
	}

}

func Calculator() {
	var x, y float32
	var operator string
	fmt.Println("Enter the first number:")
	fmt.Scan(&x)
	fmt.Println("Enter the second number:")
	fmt.Scan(&y)
	fmt.Println("Enter the operator:")
	fmt.Scan(&operator)
	ans := calculate(x, y, operator)
	fmt.Printf("The expression result is: %v", ans)
	return
}
