package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexesp/Go_Web.git/pkg/config"
	"github.com/alexesp/Go_Web.git/pkg/handlers"
	"github.com/alexesp/Go_Web.git/pkg/render"
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

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println(fmt.Sprintf("Inicio del servidor en el puerto %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
