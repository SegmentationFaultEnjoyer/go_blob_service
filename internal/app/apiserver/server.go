package apiserver

import (
	"encoding/json"
	"go_blob_service/internal/app/helpers"
	"go_blob_service/internal/app/model"
	"go_blob_service/internal/app/store"
	"net/http"
	"strconv"

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
	s.router.HandleFunc("/user", s.handleUserAdd()).Methods("POST")
	s.router.HandleFunc("/blob", s.handleBlobAdd()).Methods("POST")              //create BLOB
	s.router.HandleFunc("/blob/{id}", s.handleBlobGet()).Methods("GET")          //get BLOB by blob_ID
	s.router.HandleFunc("/blobs/{user_id}", s.handleBlobGetAll()).Methods("GET") //get all BLOBS by user_id
	s.router.HandleFunc("/blob/{id}", s.handleBlobDelete()).Methods("DELETE")    ///delete BLOB by ID
}

func (s *server) handleBlobGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		string_id := params["user_id"]

		id, err := strconv.Atoi(string_id)
		if err != nil {
			s.logger.Error("ERROR ATOI")
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		blobs, err := s.store.Blob().GetAllBlobs(id)
		if err != nil {
			s.logger.Error("ERROR GETTING ALL BLOBS")
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusAccepted, blobs)
	}
}

func (s *server) handleBlobDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		string_id := params["id"]

		id, err := strconv.Atoi(string_id)
		if err != nil {
			s.logger.Error("ERROR ATOI")
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if err := s.store.Blob().DeleteBlob(id); err != nil {
			s.logger.Error("ERROR DELETING BLOB")
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusAccepted, nil)
	}
}

func (s *server) handleBlobAdd() http.HandlerFunc {
	type request struct {
		Type        string `json:"type"`
		Title       string `json:"title"`
		LinkSelf    string `json:"link_self"`
		LinkRelated string `json:"link_related"`
		AuthorType  string `json:"author_type"`
		AuthorID    string `json:"author_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		blob, err := s.store.Blob().GenerateBlob(
			req.Type,
			req.Title,
			req.LinkSelf,
			req.LinkRelated,
			req.AuthorType,
			req.AuthorID,
		)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			s.logger.Error("ERROR GENERATING BLOB")
			return
		}

		if err := s.store.Blob().Create(blob); err != nil {
			s.logger.Error("ERROR ADDING BLOB")
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, blob)
	}
}

func (s *server) handleBlobGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		string_id := params["id"]

		id, err := strconv.Atoi(string_id)
		if err != nil {
			s.logger.Error("ERROR ATOI")
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		blob, err := s.store.Blob().GetBlobByID(id)
		if err != nil {
			s.logger.Error("ERROR GETTING BLOB")
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusFound, blob)
	}
}

func (s *server) handleUserAdd() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		user := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.store.User().Create(user); err != nil {
			s.logger.Error("ERROR ADDING USER")
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, user)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"errors": err.Error(), "status": strconv.Itoa(code)})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
