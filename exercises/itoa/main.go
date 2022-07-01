package main

import (
	"fmt"
	"strconv"
)

var arr []string

func main(){

	for i:=0; i < 30; i++{
		arr[i] = strconv.Itoa(i)
		fmt.Printf("%s %T" ,arr[i],arr[i])
	}
	
}