package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(fmt.Sprintf("hit the page: %s", req.RequestURI))
		next.ServeHTTP(res, req)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   appConfig.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// loads and save the session every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
