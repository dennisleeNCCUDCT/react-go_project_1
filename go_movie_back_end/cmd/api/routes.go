package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler{//app *application->Any variable called app can use this function(routes),and //http.Handler->means this is a http.Handler type function
	//create a route mux
	mux:=chi.NewRouter()
	mux.Use(middleware.Recoverer)//Recoverer is a middleware that recovers from panics, logs the panic (and a backtrace), and returns a HTTP 500 (Internal Server Error) status if possible. Recoverer prints a request ID if one is provided.
	
	mux.Get("/", app.Home)
	mux.Get("/movies", app.ALLMovies)
	return mux
}