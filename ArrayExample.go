package main

import "fmt"

func main() {

	var example [4][5]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			example[i][j] = i + j
		}
	}
	fmt.Println("x: ", example)

}
