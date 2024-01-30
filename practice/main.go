package main

import (
	"fmt"
)

func main() {

	type fiance struct {
		FianceName string
	}

	type Person struct {
		Name    string
		Age     int
		Company string
		Ash     fiance
	}

	kishore := Person{}
	kishore.Name = "Kishore"
	kishore.Age = 27
	kishore.Company = "Netflix"
	kishore.Ash.FianceName = "Aswathy"

	fmt.Println(kishore.Ash.FianceName)

}
