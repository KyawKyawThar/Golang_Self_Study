package main

import "fmt"

func main() {

	greeting := saySomething("Hello World")
	fmt.Println(greeting)

	myString := "White"
	fmt.Println("My String is set to", myString)

	changeUsingPointer(&myString)
	fmt.Println("After func call my String is set to", myString)
}

func saySomething(s string) string {
	return s
}

func changeUsingPointer(p *string) {
	//fmt.Println("Pointer", p)
	//fmt.Println("dereference", &p)
	color := "Blue"
	*p = color
}
