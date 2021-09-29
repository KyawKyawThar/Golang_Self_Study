package main

import (
	"fmt"
	"sort"
)

func main() {
	var mySlice []int

	mySlice = append(mySlice,2)
	mySlice = append(mySlice,1)
	mySlice = append(mySlice,3)

	sort.Ints(mySlice)

	fmt.Println(mySlice)


	number := []int {3,5,1,5,8,3,6,7,6,4}
	fmt.Println(number[:5])
	fmt.Println(number[4:])
	fmt.Println(number[3:8])
}
