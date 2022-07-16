package sqlstore

import (
	"database/sql"
	"go_blob_service/internal/app/helpers"
	"go_blob_service/internal/app/store"

	_ "github.com/lib/pq"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		logger, _ := helpers.ConfigureLogger("debug")

		s.userRepository = &UserRepository{
			store:  s,
			logger: logger,
		}
	}

	return s.userRepository
}
