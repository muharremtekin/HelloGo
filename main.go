package main

import (
	"fmt"
	"learn-go/essentials"
)

func main() {
	fmt.Println("Hello World!")
	// essentials.DiffrentVariables()
	// fmt.Println("Base URL:", essentials.BASE_URL)

	// essentials.Slices()
	// essentials.Arrays()
	// essentials.Structs()
	essentials.Maps()
	// wait for user input before exiting
	var input string
	fmt.Println("Press Enter to exit...")
	fmt.Scanln(&input)
}
