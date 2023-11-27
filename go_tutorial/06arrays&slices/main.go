package main

import "fmt"

func main()  {

	//Arrays
	// var ages[3]int= [3]int{20,25,30}
	var ages = [3]int{20,25,30}
	fmt.Println(ages,len(ages))

	names:=[3]string{"Kishore","Aswathy","Ishan"}
	fmt.Println(names,len(names))

	//Slices {uses Arrays under the hood}

	var scores = [] int {10,43,55,66}
	scores=append(scores,100)
	fmt.Println(scores,len(scores))

	//slice ranges 
	rangeOne:=scores[1:4]
	rangeTwo:=scores[1:]
	rangeThree:=scores[:4]
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)


	
}