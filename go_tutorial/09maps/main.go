package main

import "fmt"

func main()  {
	menu:=map[string]float32{
		"soup":4.99,
		"pie":7.99,
		"salad":8.44,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	menu["salad"]=10.5

	//looping maps
	for k,v:=range menu{
		fmt.Println(k,v)
	}
}