package main

import "fmt"

func main() {
	const name string = "Kishore"
	fmt.Println("Hello", name)
	var fruits []string
	fruits = append(fruits, "Mango")
	fruits = append(fruits, "Apple")
	fruits = append(fruits, "Orange")
	fruits = append(fruits, "Grapes")
	fruits = append(fruits, "Strawberry")

	fmt.Println(fruits)
}
