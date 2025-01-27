package handlers

import (
	"net/http"

	"github.com/alexesp/Go_Web.git/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Es la pagina de inicio")
	render.RenderTemplate(w, "home.page.tmpl")
}
func About(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Es la pagina sobre!")
	render.RenderTemplate(w, "about.page.tmpl")
}
