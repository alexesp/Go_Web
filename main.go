package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

}
func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	_, err := fmt.Fprintf(w, "Hello Web!")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// })
	http.HandleFunc("/", Home)
	_ = http.ListenAndServe(":8088", nil)
}
