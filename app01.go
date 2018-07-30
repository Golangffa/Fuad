package main

import "fmt"

func main() {
	//var a int
	//var is introducing a variable
	a := 5
	var b int = 2
	//var c int
	//fmt.Printf("%d\n", a / b)
	//%d is used to display an integer
	if a%b == 0 {
		fmt.Println("a is even")
	} else {
		fmt.Println("a is odd")
	}
	//a%b is the remainder of a divide b
	//== means is equal to
	if (a+b > 100) || (a-b > 100) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	fmt.Println("Can I have a loan of £200?")

	if (50 > 100) && (1000 > 100) {
		//&& is and; if one is neg and one is pos the outcome will be neg
		//|| is or; if one is neg and one is pos the outcome will be pos
		fmt.Println("Yes, you have enough dosh")
	} else {
		fmt.Println("Bog off")
	}
	fmt.Println("Can i have a small loan of one grand?")

	if 5*52*5 >= 1000 {
		fmt.Println("Yes i will grant you a small loan of 1000")
	} else {
		fmt.Println("Get a job")
	}

	sum := 0
	for i := 0; i < 10; i++ {
		//the for loop contains init statement = what i is equal to
		//the condition expression = what i has to be
		//post statement = +1
		//the loop will go around until i is nine
		fmt.Printf("This is i: %d\n", i)
		sum += i
		//sum += i states to value of sum
		fmt.Printf("This is sum: %d\n", sum)
		//Printf will  always need %x depending on

	}
	fmt.Println(sum)
}
