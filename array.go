package main

import "fmt"

func LearnArray() {
	//arr := [5]int{1, 2, 3, 4}
	//arr[1] = 7

	var arr = [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	return
}
