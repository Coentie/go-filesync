package router

import (
	"github.com/coentie/filesync-server/controllers"
	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", controllers.Upload)
	r.Get("/files", controllers.Files)
	r.Get("/download/{filename}", controllers.Download)

	return r
}
