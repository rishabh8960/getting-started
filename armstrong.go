//check a number is armstrong or not
package main

import "fmt"

func main() {
	var a, b, c, d, e, f int
	fmt.Print("enter a number to check armtrong : ")
	fmt.Scanln(&a)
	e = 0
	for c = a; c > 0; c = c / 10 {
		f = c % 10
		b = 1
		for d = a; d > 0; d = d / 10 {
			b = b * f
		}
		e = e + b
	}
	if e == a {
		fmt.Print("number is armstrong")
	} else {
		fmt.Print("number is not armstrong")
	}
}
