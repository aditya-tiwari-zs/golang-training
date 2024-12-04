package assignment1

import "fmt"

func circlePerimeter(side float32) float32 {
	const pi = 3.14
	ans := 2 * pi * side
	return ans
}

func squarePermimeter(side float32) float32 {
	ans := 4 * side
	return ans
}

func Perimeter() {

	var side float32

	fmt.Println("Please give side length:")
	fmt.Scan(&side)
	circle := circlePerimeter(side)
	square := squarePermimeter(side)

	fmt.Printf("Circle perimeter:%v \n", circle)
	fmt.Printf("Circle perimeter:%v", square)

}
