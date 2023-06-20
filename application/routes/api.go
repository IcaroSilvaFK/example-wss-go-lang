package routes

import (
	"net/http"

	"github.com/IcaroSilvaFK/example-wss-go-lang/application/controllers"
	"github.com/go-chi/chi"
)

func InitializeApiRoutes(mux *chi.Mux) {

	mux.Get("/heath", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux.Get("/room", controllers.NewWssController)

}
