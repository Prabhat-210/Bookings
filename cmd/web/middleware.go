package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { //func here is anonyms function which is casted to handlerFunc
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r) //next will pass to next page
	})
}

// adds csrf protection to all post request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next) //it will create a handler for us
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProducation,
		SameSite: http.SameSiteLaxMode,
	}) //We need to set the base cookie because it uses cookies to make sure that the token it generates is available on a per page basis.
	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
