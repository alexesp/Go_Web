package main

import (
	"net/http"
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
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	_ = http.ListenAndServe(portNumber, nil)
}
