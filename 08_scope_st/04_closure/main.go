package main

import "fmt"

func main() {
	//  inside scope of (x)
	x := "scope of x"
	fmt.Printf("This is %s. \n", x)
	{
		y := "scope of Y"
		fmt.Printf("And this is %s inside %s.", x, y)
	}
}
