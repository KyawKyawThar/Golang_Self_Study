package main

import (
	"errors"
	"fmt"
)

func main() {

	myNum, err := divided(100, 10.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result of division is", myNum)
}

func divided(x int16, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("cannot divided by 0")
	}

	result = float32(x) / y

	return result, nil
}
