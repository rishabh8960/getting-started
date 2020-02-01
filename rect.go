package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("enter hight of rectangle")
	fmt.Scanln(&a)
	fmt.Println("enter width of rectangle")
	fmt.Scanln(&b)
	fmt.Println("area of rectangle is :", (a * b), "\nperimeter of rectangle is : ", (2 * (a + b)))
}
