package slice

import "fmt"

func swap(slice *[]int, i, j int) {
	temp := (*slice)[i]
	(*slice)[i] = (*slice)[j]
	(*slice)[j] = temp
}

func reverse(slice *[]int) {
	i := 0
	j := len(*slice) - 1

	for i <= j {
		swap(slice, i, j)
		i++
		j--
	}

	fmt.Println(*slice)
}

func LearnSlice() {
	slice := []int{1, 2, 3, 4, 5, 6}

	reverse(&slice)

}
