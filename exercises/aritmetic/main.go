package main

import "fmt"

var (
	a float64 = 3.4
	b float32 = 8.6
	c int16 = 35
	e uint16 = 7
)

func main(){

	result1 := a - float64(b)

	fmt.Printf("a is %T, b is %T. Result of a - float64(b) is %g\n",a,b,result1)

	result2 := c + int16(b)

	fmt.Printf("c is %T, b is %T. Result of c + int16(b) is %d\n",c,b,result2)

	result3 := uint16(c) / e

	fmt.Printf("c is %T, e is %T. Result of uint16(c) is %d\n",c,e,result3)

	result4 := cube(a)

	fmt.Printf("a is %T, result4 is %T. Result is %g\n",a,result4,result4)
}

func cube(number float64) float64{
	for i:=0; i<1; i++{
		number *= number
	}
	return number
}