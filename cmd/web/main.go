package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gobooking/pkg/config"
	"gobooking/pkg/handlers"
	"gobooking/pkg/render"

	"github.com/alexedwards/scs/v2"
)

var appConfig config.AppConfig
var session *scs.SessionManager

func main() {
	appConfig.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.InProduction

	appConfig.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	appConfig.TemplateCache = templateCache
	appConfig.UseCache = false
	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)
	render.NewTemplates(&appConfig)

	srv := &http.Server{
		Addr:    ":8282",
		Handler: routes(&appConfig),
	}
	fmt.Println(fmt.Sprintf("http://localhost:8282"))
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
