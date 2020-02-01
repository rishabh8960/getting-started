//linearsearch
package main

import "fmt"

func main() {
	var n int
	fmt.Print("enter the number of element you want to enter : ")
	fmt.Scanln(&n)
	var a [100]int
	var b, c int
	c = 0
	fmt.Print("enter the  elements\n")
	for i := 0; i < n; i++ {
		fmt.Scanln(&a[i])
	}
	fmt.Print("enter the element you want to search in list : ")
	fmt.Scanln(&b)
	for i := 0; i < n; i++ {
		if a[i] == b {
			fmt.Print("Element '", a[i], "' is present in the list\nIndex of Element is : ", i)
			c++
			break
		}
	}
	if c == 0 {
		fmt.Print("Element '", b, "' is not present in the list")

	}
}
