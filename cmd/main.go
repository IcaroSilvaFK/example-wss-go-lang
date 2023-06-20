package main

import (
	"net/http"

	"github.com/IcaroSilvaFK/example-wss-go-lang/application/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	mux := chi.NewRouter()

	mux.Use(middleware.Logger)

	routes.InitializeApiRoutes(mux)
	http.ListenAndServe(":8000", mux)
}
