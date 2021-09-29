package main

import "fmt"

type User struct {
	firstName string
	lastName  string
}

func main() {
	myMap := make(map[string]int)
	myMap["Ferrari"] = 3
	myMap["Lambo"] = 8
	myMap["Ferrari"] = 15

	fmt.Println(myMap["Ferrari"])

	myMap1 := make(map[string]User) //User struct type is important
	me := User{firstName: "Highest", lastName: "LeveL"}
	myMap1["me"] = me

	fmt.Println(myMap1["me"].firstName, myMap1["me"].lastName)

}
