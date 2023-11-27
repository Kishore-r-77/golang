package main

import "fmt"

func updateName(n *string)  {
	*n="Aswathy"
}

func main()  {
	
	name:="Kishore"
	m:=&name
	fmt.Println(m)
	fmt.Println(*m)
	updateName(m)
	fmt.Println(name)
}