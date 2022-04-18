package main

import (
	"github.com/okiprakasa/hello-world/pkg/config"
	"github.com/okiprakasa/hello-world/pkg/handlers"
	"github.com/okiprakasa/hello-world/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.TemplatePointer(&app, err)

	srv := &http.Server{ //serve
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
