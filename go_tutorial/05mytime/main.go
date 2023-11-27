package main

import (
	"fmt"
	"time"
)

func main()  {
	presentTime:=time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate:=time.Date(2022,time.April,21,23,30,0,0,time.UTC)

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}