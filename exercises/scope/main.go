package main

import "fmt"

func main(){

	fmt.Println("End of the result of scope()/smt is",scope()) 

	fmt.Println("End of the result of scope2()/smt is",scope2())
}

func scope() int{				// I declared again smt  for all of the layer.
	smt := 30
	fmt.Println("smt in 1st layer is",smt)
	for smt:=15; smt>9; smt--{
		if smt == 13{
			smt := 11
			fmt.Println("smt in 3rd layer is",smt)
		}
		fmt.Println("smt in 2nd layer is",smt)
	}
	return smt
}

// OUTPUT;
// smt in 1st layer is 30
// smt in 2nd layer is 15
// smt in 2nd layer is 14
// smt in 3rd layer is 11
// smt in 2nd layer is 13
// smt in 2nd layer is 12
// smt in 2nd layer is 11
// smt in 2nd layer is 10
// End of the result of scope()/smt is 30

func scope2() int{				// I declared smt at once then change its values.
	smt := 30
	fmt.Println("smt in 1st layer",smt)
	for smt=15; smt>9; smt--{
		if smt == 13{
			smt = 11
			fmt.Println("smt in 3rd layer",smt)
		}
		fmt.Println("smt in 2nd layer",smt)
	}
	return smt
}

// OUTPUT :
// smt in 1st layer 30
// smt in 2nd layer 15
// smt in 2nd layer 14
// smt in 3rd layer 11
// smt in 2nd layer 11
// smt in 2nd layer 10
// End of the result of scope2()/smt is 9