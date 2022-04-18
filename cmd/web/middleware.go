package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

//WriteToConsole is a middleware used to print to console everytime page is asked
func _(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Page loaded")
		next.ServeHTTP(w, r)
	})
}

//NoSurf creates CSRF Token for every page, used when submitting form as a hidden value
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false, //Not making https
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
