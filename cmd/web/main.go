package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/okiprakasa/hello-world/pkg/config"
	"github.com/okiprakasa/hello-world/pkg/handlers"
	"github.com/okiprakasa/hello-world/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

//main is the main application function
func main() {
	//Change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 6 * time.Hour //Available for 6 hours of idle / no page request
	session.Cookie.Persist = true    //Persist even if the web browser window is closed
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction //Localhost is not https secure (Development Mode)

	//Casting the session var to AppConfig Session Parameter
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc

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
