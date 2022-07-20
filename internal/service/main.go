package service

import (
	"fmt"
	"net"
	"net/http"

	"testService/internal/config"
	"testService/internal/store"
	"testService/internal/store/sqlstore"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
	store    store.Store
}

func (s *service) run() error {

	s.log.Info("Service started")
	r := s.router()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}
	fmt.Println("SERVICE STARTED")
	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {

	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
		store:    sqlstore.New(cfg.DB().RawDB()),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
