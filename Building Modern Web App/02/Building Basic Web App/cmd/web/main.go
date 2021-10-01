package main

import (
	"fmt"
	"github.com/hl/basic-web/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Server is running on port %s ", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
