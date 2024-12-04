package assignment1

import "fmt"

func InputString() {
	var name string
	fmt.Println("What is your name ?")
	fmt.Scan(&name)
	fmt.Printf("Hello %s \n", name)
}
