package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string,string) {
	s:=strings.ToUpper(n)
	splittedValues:=strings.Split(s, " ")


	var initials []string
	for _,v :=range splittedValues{
		initials=append(initials,v[:1])
	}
	if(len(initials)>1){
		return initials[0],initials[1]
	}
	return initials[0],"_"
}

func main()  {
	fn,sn:=getInitials("Kishore")
	fmt.Println(fn,"-",sn)
}