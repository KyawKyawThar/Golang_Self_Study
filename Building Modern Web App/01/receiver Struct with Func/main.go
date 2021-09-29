package main

import "fmt"

type myStruct struct {
	firstName string
}

func (s *myStruct) receiver() string {
	return s.firstName
}

func main() {
	var myVar myStruct
	myVar.firstName = "Kyaw Kyaw"

	myVar2 := myStruct{firstName: "HighestLeveL"}

	fmt.Println('R',myVar.receiver())
	fmt.Println('R',myVar2.receiver())
	fmt.Println(myVar.firstName)
	fmt.Println(myVar2.firstName)
}
