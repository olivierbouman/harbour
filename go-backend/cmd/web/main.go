package main

import (
	"fmt"
	"go-backend/pkg/config"
	"go-backend/pkg/handlers"
	"go-backend/pkg/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// Main is the main applications function
func main() {
	//Change this to true when in production
	app.InProduction = false

	// Refering to variable that exists outside the scope of main()
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// Create your templateCache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create templateCache")
	}

	// Store our tc in out app variable from appConfig
	app.TemplateCache = tc
	app.UseCache = false

	// Creating our repo
	repo := handlers.NewRepo(&app)

	// And passing is back to the handlers
	// This way we can pass down our AppConfig to be used in handlers
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println((fmt.Sprintf("Starting application on port %s", portNumber)))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: chi_router(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
