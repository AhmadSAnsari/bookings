package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AhmadSAnsari/bookings/pkg/config"
	"github.com/AhmadSAnsari/bookings/pkg/handlers"
	"github.com/AhmadSAnsari/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager


// main is the main application function
func main() {

	// Change this to true when in Prodiction
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour	//for 24 hours
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode

	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println("Starting application on port:", portNumber)

	srv := &http.Server{
		Addr : portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
