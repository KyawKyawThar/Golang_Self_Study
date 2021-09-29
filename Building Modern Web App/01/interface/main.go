package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {

	dog := Dog{
		Name:  "Jackei",
		Breed: "Mammel",
	}
	printInfo(dog)

	gorilla := Gorilla{Name: "King Kong", Color: "Black", NumberOfTeeth: 43}
	printInfo(gorilla)
}

func (d Dog) Says() string {
	return "Woolf"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

func (g Gorilla) Says() string {
	return "Garr & Black"
}

func (g Gorilla) NumberOfLegs() int {
	return 2
}

func printInfo(a Animal) {
	fmt.Println("This animals says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}
