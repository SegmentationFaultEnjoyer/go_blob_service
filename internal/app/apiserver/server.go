package apiserver

import (
	"go_blob_service/internal/app/helpers"
	"go_blob_service/internal/app/model"
	"go_blob_service/internal/app/store"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: nil,
		store:  store,
	}

	s.configureLogger()
	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureLogger() error {
	logger, err := helpers.ConfigureLogger("debug")

	if err != nil {
		return err
	}

	s.logger = logger

	return nil
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/add", s.handleAdd())
}

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}

func (s *server) handleAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := s.TestDB("as", "aaaa"); err != nil {
			s.logger.Error("ERROR ADDING USER")
			io.WriteString(w, "ERROR ADDING")
			return
		}

		s.logger.Info("USER HAS BEEN ADD")
		io.WriteString(w, "USER HAS BEEN ADD")
	}
}

func (s *server) TestDB(email string, password string) error {
	if err := s.store.User().Create(
		&model.User{
			Email:             email,
			EncryptedPassword: password,
		}); err != nil {
		return err
	}

	return nil
}
