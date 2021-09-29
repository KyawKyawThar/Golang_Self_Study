package main

import (
	"encoding/json"
	"fmt"
)

var myJson = []byte(`[
 {
   "FirstName" : "Platyus",
   "LastName" : "Aeidad",
   "HairColor" : "Red",
   "HasDog" : "true"
},{

  "FirstName" : "Quoll",
   "LastName" : "Dasyu",  
   "HairColor" : "Red",
   "HasDog" : "false"
}
]`)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	var unmarshalled []Person

	err := json.Unmarshal(myJson, &unmarshalled)

	if err != nil {
		fmt.Println("Error Unmarshalling json: ", err)
	}

	fmt.Printf("%+v", unmarshalled)

	var mySlice []Person

	var m1 Person
	m1.FirstName = "Htoo Thiri"
	m1.LastName = "Shein"
	m1.HairColor = "Black"
	m1.HasDog = true

	var m2 Person
	m2.FirstName = "Zar Zar"
	m2.LastName = "Htun"
	m2.HairColor = "Gold"
	m2.HasDog = false

	mySlice = append(mySlice, m1, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(newJson))
}
