package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Es la pagina de inicio")
	renderTemplate(w, "home.page.tmpl")
}
func About(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Es la pagina sobre!")
	renderTemplate(w, "about.page.tmpl")
}
