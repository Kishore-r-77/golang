package main

import (
	"fmt"
)

func main() {
	input := "John 25"
	var name string
	var age int

	n, err := fmt.Sscanf(input, "%s %d", &name, &age)

	fmt.Println(n)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Scanned %d items. Name: %s, Age: %d\n", n, name, age)
	}
}
