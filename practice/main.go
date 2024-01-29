package main

import "fmt"

func test(num *int) {
	*num = 7
}
func main() {

	const name string = "Kishore"
	var num int = 5

	fmt.Println(name)
	fmt.Printf("The number is %.2f\n", 10.0)
	test(&num)
	fmt.Println(num)

}
