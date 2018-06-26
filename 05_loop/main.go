package main

import "fmt"

func main() {
	for i := 10000000; i < 100000000; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}
