package main

import (
	"fmt"
	"os"
)

func main() {
	argLength := len(os.Args[1:])
	if argLength > 0 {
		arg := os.Args[1]
		fmt.Println(arg)
	} else {
		fmt.Println("Hello World")
	}
}
