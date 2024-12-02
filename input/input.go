package input

import "fmt"

func TakeInput() {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scan(&name)
	fmt.Printf("Your given name is %s", name)
}
