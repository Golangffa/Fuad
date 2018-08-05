package main

import "fmt"

func main() {
	x := make(map[string]int)

	x["s1"] = 6
	x["s2"] = 12

	v1 := x["s1"]
	v2 := x["s2"]
	fmt.Println("examplemap:", v2-v1)

}
