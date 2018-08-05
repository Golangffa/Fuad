package main

import "fmt"

func main() {

	x := make([]string, 5)
	fmt.Println("emp:", x)

	x[0] = "l"
	x[1] = "m"
	x[2] = "n"
	x[3] = "o"
	x[4] = "p"
	a := x[1:4]
	fmt.Println("slice:", a)

}
