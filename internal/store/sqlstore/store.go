package sqlstore

import (
	"database/sql"
	"testService/internal/service/helpers"
	"testService/internal/store"

	_ "github.com/lib/pq"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
	blobRepository *BlobRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Blob() store.BlobRepository {
	if s.blobRepository == nil {
		logger, _ := helpers.ConfigureLogger("debug")
		s.blobRepository = &BlobRepository{
			store:  s,
			logger: logger,
		}
	}

	return s.blobRepository
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
