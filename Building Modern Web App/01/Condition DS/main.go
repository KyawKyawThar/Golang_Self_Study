package main

import "fmt"

func main() {
	animal := "dog"

	switch animal {
	case "dog":
		fmt.Println("This is a dog")
	case "cat":
		fmt.Println("This is a cat")

	case "fish":
		fmt.Println("This is a fish")

	default:
		fmt.Println("Dog is something else")
	}
}
