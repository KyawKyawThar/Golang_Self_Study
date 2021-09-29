package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

//Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is home page")
}

//About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValue(2, 5)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 5 is %d", sum))
}

//Divide is the divided value
func Divide(w http.ResponseWriter, r *http.Request) {
	d, err := didivedValue(100.0, 00.0)
	if err != nil {
		fmt.Fprintf(w, "Can not divided by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided %f is %f", 100.0, 00.0, d))
}

//AddValue adds two integers and return the sum.
func addValue(x, y int) int {
	return x + y
}

func didivedValue(x, y float32) (float32, error) {

	if y <= 0 {
		err := errors.New("Cannot divided by zero")
		return 0, err
	}

	result := x / y

	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Server is running on port %s ", portNumber))
	log.Fatalln(http.ListenAndServe(portNumber, nil))
}
