// This code convert Kilometer to Meter or Centimeter;
// Meter to KM or CM;
// Centimeter to M or KM.

// Kilometer	km	10^3
// Meter		m	1
// Centimeter	cm	10^-2

package main

import (
	"fmt"
	"os"
)

type (
	Meter float64
	Kilometer float64
	Centimeter float64
)

var (
	a Meter = 1
	b Kilometer = 1
	c Centimeter = 1
)

func main(){

	word := os.Args[1]		// You should enter word like ktom,ktoc,mtok,mtoc,ctok,ctom to convert. 

	if word == "ktom" {  // Kilometer to Meter
		fmt.Println(KtoM(b))
	}else if word == "ktoc" { // Kilometer to CM
		fmt.Println(KtoC(b))					
	}else if word == "mtok" {	//Meter to KM
		fmt.Println(MtoK(a))
	}else if word == "mtoc" {	// Meter to CM
		fmt.Println(MtoC(a))
	}else if word == "ctok" {	// CM to KM
		fmt.Println(CtoK(c))
	}else if word == "ctom" {	// CM to Meter
		fmt.Println(CtoM(c))
	}

}

func KtoM(k Kilometer) Meter { return Meter(k)*10*10*10}
func KtoC(k Kilometer) Centimeter { return Centimeter (k)*10*10*10*10*10}

func MtoK(m Meter) Kilometer { return Kilometer(m)/10/10/10}
func MtoC(m Meter) Centimeter { return Centimeter(m)*10*10}

func CtoK(m Centimeter) Kilometer { return Kilometer(m)/10/10/10/10/10}
func CtoM(m Centimeter) Meter{ return Meter(m)/10/10}
