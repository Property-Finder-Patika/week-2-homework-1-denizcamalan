package main

import "fmt"

func main() {

	hello() // Call the bye function from inside the hello function

}

func bye(){
	fmt.Print("bye bye")
}
