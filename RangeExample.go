package main

import "fmt"

func main() {

	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for _, num := range num {
		sum += num
	}
	fmt.Println("sum:", sum)

}
