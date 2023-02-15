package cmd

import (
	"os"
)

func main() {
	argLength := len(os.Args[1:])
	if argLength > 0 {
		cityName := os.Args[1]
		Run(cityName)
	} else {
		Run("")
	}
}
