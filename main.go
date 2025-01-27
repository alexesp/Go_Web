package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8088"

func Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Es la pagina de inicio")
	renderTemplate(w, "home.page.tmpl")
}
func About(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Es la pagina sobre!")
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsin tempate:", err)
		return
	}
}
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
