package main

import "fmt"

func main() {
	var a int
	var b float64
	fmt.Print("enter your choice\n1 for rupee to dollar\n2.for dollar to rupee\n")
	fmt.Scanln(&a)
	fmt.Print("enter ammount : ")
	fmt.Scanln(&b)
	switch a {
	case 1:
		fmt.Print("ammount in dollar is : ", b/70)
	case 2:
		fmt.Print("ammount in rupee is : ", b*70)
	}
}
