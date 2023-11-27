package main

import "fmt"

const LoginToken string="Uxkfkdslk"

func main(){
	var username string="Kishore"
	fmt.Printf("Variable is of type: %T \n",username)

	var isLogged bool=true
	fmt.Printf("Variable is of type: %T \n",isLogged)

	var smallVal uint8=255
	fmt.Printf("Variable is of type: %T \n",smallVal)

	var smallFloat float32=255.5565675757675
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n",smallFloat)


	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)

	//implicit type
	var website = "google.com"
	fmt.Println(website)

	//no var style
	numbOfUser:=7
	fmt.Println(numbOfUser)

	fmt.Println(LoginToken)

	fmt.Printf("My name is %s","yukesh")
}