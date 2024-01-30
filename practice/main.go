package main

import "fmt"

func main() {
	var name string = "Aswathy"

	switch name {
	case "Kishore":
		{
			fmt.Println("Hello Kishore")
		}
	case "Aswathy":
		fmt.Println("Hello Aswathy")
		fmt.Println("Hello Goddess")
	default:
		fmt.Println("Select a Name")
	}

}
