package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("Welcome to golang")

	content:="This needs to go in a file - www.kishore.com"

	file,err := os.Create("./myfile.txt")
	if err!=nil{
		panic(err)
	}

	length,err:=io.WriteString(file,content)

	if err!=nil{
		panic(err)
	}

	fmt.Println("The length is: ",length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string){
	databyte,err:=ioutil.ReadFile(filename)
	if err!=nil{
		panic(err)
	}
	fmt.Println("Text data inside the file is \n",string(databyte))

}







