package main

import (
	"fmt"
	"github.com/KyawKyawThar/myniceprogram/helpers"
)

func main() {

	var myVar helpers.SomeType
	myVar.TypeName = "HL"
	myVar.TypeNumber = 14
	fmt.Println(myVar.TypeNumber)
	fmt.Println(myVar.TypeName)
}

