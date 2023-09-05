package main

import (
	"fmt"
)

func main() {
	var str string = "This is a string in Go" // Notice the type is explicity defined.
	fmt.Println(str)
	fmt.Println("Lets see how this works")
	// There can be only 1 main function in Go package/module.
	// Go supports pointers and functions can be passed around as parameters to other functions.
	// Go is statically typed language and you will have to specify the type of each varible during runtime.
	var i int = 10
	fmt.Printf("The type of variable is %T\n", i)

	var another_string = "This is a string in Go"
	fmt.Printf("The type of variable is %T\n", another_string) // compiler figures out the type.

	myString := "my string" // when you use colon and equals, no need to use the var keyword.
	fmt.Printf("The type of variable is %T\n", myString)

}
