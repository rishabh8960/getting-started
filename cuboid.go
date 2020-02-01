//calculate surface area and volume of a cuboid
package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("enter length of cuboid")
	fmt.Scanln(&a)
	fmt.Println("enter width of cuboid")
	fmt.Scanln(&b)
	fmt.Println("enter hight of cuboid")
	fmt.Scanln(&c)
	fmt.Println("volume of cuboid is : ", (a * b * c), "\narea of cuboid is : ", (2 * ((a * b) + (b * c) + (c * a))))
}
