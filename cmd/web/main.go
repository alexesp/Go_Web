package main

import (
	"fmt"
	"net/http"

	"github.com/alexesp/Go_Web.git/pkg/handlers"
)

const portNumber = ":8088"

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	_, err := fmt.Fprintf(w, "Hello Web!")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// })
	//fmt.Println("Test")
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println(fmt.Sprintf("Inicio del servidor en el puerto %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
