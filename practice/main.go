package main

import "fmt"

func main() {

	var name string

	fmt.Println("What is you're name")

	fmt.Scan(&name)

	fmt.Printf("Hello %v", name)

}
