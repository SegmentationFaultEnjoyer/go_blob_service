package service

import (
	"testService/internal/service/handlers"
	"testService/internal/service/requests"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	controller := requests.NewController(s.store)

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
		),
	)
	r.Route("/integrations/testService", func(r chi.Router) {
		// configure endpoints here
		r.Post("/user", controller.HandleUserAdd())
		r.Post("/blob", controller.HandleBlobAdd())
		r.Get("/blob/{id}", controller.HandleBlobGet())
		r.Get("/blobs/{user_id}", controller.HandleBlobGetAll())
		r.Delete("/blob/{id}", controller.HandleBlobDelete())
	})

	return r
}
