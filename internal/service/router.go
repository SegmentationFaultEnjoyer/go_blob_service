package service

import (
	"testService/internal/service/handlers"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxDB(s.db),
		),
	)
	r.Route("/integrations/testService", func(r chi.Router) {
		// configure endpoints here
		r.Post("/blob", handlers.HandleBlobAdd)
		r.Get("/blob/{id}", handlers.HandleBlobGet)
		r.Get("/blob", handlers.HandleBlobGetAll)
		r.Delete("/blob/{id}", handlers.HandleBlobDelete)
	})

	return r
}
