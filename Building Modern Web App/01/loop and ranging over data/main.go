package main

import "fmt"

type User struct {
	firName  string
	lastName string
}

func main() {
	//mySlice := []string{"cat","dog","fish","birds","horse"}
	//
	//for i, x := range mySlice {
	//	fmt.Println(i,x)
	//}

	//myMap := make(map[string]string)
	//
	//myMap["dog"] = "dog"
	//myMap["cat"] = "cat"
	//myMap["fish"] = "fish"
	//
	//for i, x := range myMap {
	//	fmt.Println(i, x)
	//}

	var mySlice []User

	u1 := User{
		firName:  "Highest",
		lastName: "LeveL",
	}

	u2 := User{
		firName:  "Kyaw Kyaw",
		lastName: "Thar",
	}

	mySlice = append(mySlice,u1)
	mySlice = append(mySlice,u2)

	for i,x := range mySlice {
		fmt.Println(i,x.firName)
	}
}
