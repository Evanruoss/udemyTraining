package main

import (
	"fmt"

	"github.com/evanruoss/udemyTraining/08_scope_st/03_visibility/visvars"
)

func main() {
	fmt.Println(visvars.MyName)
	visvars.PrintVar()
}

//	func main prints one println(MyName)
//	MyName located in ext pkg (visvars) at package level scope
