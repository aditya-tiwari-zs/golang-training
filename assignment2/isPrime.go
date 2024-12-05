package assignment2

import "fmt"

func isPrime(number int) bool {
	var count int = 0
	for i := 1; i < number; i++ {
		if number%i == 0 {
			count++
		}
	}

	if count > 1 {
		return false
	}
	return true
}

func FindPrime() {
	var number int
	fmt.Println("Enter your number:")
	fmt.Scan(&number)
	ans := isPrime(number)
	if ans {
		fmt.Println("Number is prime")
		return
	}
	fmt.Println("Number is not a prime")
	return
}
