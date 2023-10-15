package main

import (
	"fmt"
	"github.com/Prabhat-210/Bookings/pkg/config"
	"github.com/Prabhat-210/Bookings/pkg/handler"
	"github.com/Prabhat-210/Bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager //define it here so that other file access session

func main() {

	app.InProducation = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour //till what time session will stay open

	//as all sessions using cookie so we need to set some parameters for that session
	session.Cookie.Persist = true //if true then cookie will be persist in browser
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProducation

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("can't create template cache")
	}

	app.TemplateCache = tc

	app.UseCache = false // set false for developer mode and save true for production

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplates(&app) // call will pass the template to render.go

	//http.HandleFunc("/", handler.Repo.Home)
	//http.HandleFunc("/about", handler.Repo.About)
	fmt.Println(fmt.Sprintf("Starting Port at port number %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
