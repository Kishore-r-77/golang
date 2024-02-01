package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

type Animal struct {
	name string
	age  int
}

type Display interface {
	makeNoise()
}

func (person Person) makeNoise() {
	fmt.Println("Hey Yooooo")
}
func (animal Animal) makeNoise() {
	fmt.Println("Roarrrrr")
}

func display(d Display) {
	d.makeNoise()
}

func main() {

	kishore := Person{}
	kishore.name = "Kishore"
	kishore.age = 27

	tiger := Animal{}
	tiger.name = "Tiger"
	tiger.age = 20

	display(kishore)

}
