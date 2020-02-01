package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("enter an number : ")
	fmt.Scanln(&a)
	c = 0
	for b = a; b > 0; b = b / 10 {
		c = c + (b % 10)
	}
	fmt.Print("SUM IS : ", c)
}
