package main

import (
	"fmt"
	"os"
	"strconv"
)


func main(){

	msg := os.Args[1]

	c := StoI(msg)
	b := int64(c)

	msg = strconv.FormatInt(b,6)
	
	fmt.Printf("type is %T, value is at base 6 %d ",StoI(msg),StoI(msg))
}

func StoI(s string) int{
	result ,err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}
	return result
}