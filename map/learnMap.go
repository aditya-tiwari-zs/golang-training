package _map

import "fmt"

func wordMatrix(word string) {
	data := make(map[string]int)
	for _, val := range word {
		data[string(val)]++
	}

	for key, val := range data {
		fmt.Printf("%s : %d\n", key, val)
	}
}

func sumValuesByKey(stringMap map[string][]int) {
	updatedMap := make(map[string]int)
	for key, value := range stringMap {
		sum := 0
		for _, data := range value {
			sum = sum + data
		}
		updatedMap[key] = sum
	}

	for key, val := range updatedMap {
		fmt.Printf("%s : %d\n", key, val)
	}
}

func LearnMap() {
	//word := "geeksforgeek"

	stringMap := map[string][]int{"aditya": {1, 2, 3, 4}, "tiwari": {2, 3, 4, 5}}

	//wordMatrix(word)

	sumValuesByKey(stringMap)
}
