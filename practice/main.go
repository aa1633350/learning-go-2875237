package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin) // bufio package is used for I/O operations.
	fmt.Print("Enter your text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered string", input)

	fmt.Print("Enter your text: ")

}
