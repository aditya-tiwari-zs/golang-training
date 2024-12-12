package main

import "fmt"

func Sort(arr []int, comparator func(a, b int) bool) {
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if comparator(arr[i], arr[j]) {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	fmt.Println(arr)
}

func main() {
	arr := []int{3, 4, 2}

	comparator := func(a, b int) bool {
		return a > b
	}

	Sort(arr, comparator)
}
