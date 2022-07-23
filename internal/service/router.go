package service

import (
	"testService/internal/service/handlers"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	controller := handlers.NewController(s.db)

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
		),
	)
	r.Route("/integrations/testService", func(r chi.Router) {
		// configure endpoints here
		r.Post("/blob", controller.HandleBlobAdd())
		r.Get("/blob/{id}", controller.HandleBlobGet())
		r.Get("/blob", controller.HandleBlobGetAll())
		r.Delete("/blob/{id}", controller.HandleBlobDelete())
	})

	return r
}
