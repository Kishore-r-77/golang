package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	fmt.Println("user input")

	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")

	//comma ok || error ok
	input,_:=reader.ReadString('\n')

	

	fmt.Println("Thanks for rating: ",input)
	fmt.Printf("Type of this input is %T: ",input)
}