package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

//WriteToConsole is a middleware used to print to console everytime page is asked
//(example as reference to other func)
func _(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Page loaded")
		next.ServeHTTP(w, r)
	})
}

//NoSurf creates CSRF Token for protection to all POST request, used when submitting form as a hidden value
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",              //Applied to all pages
		Secure:   app.InProduction, //Not making https
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

//SessionLoad loads and saves the session on every page request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
