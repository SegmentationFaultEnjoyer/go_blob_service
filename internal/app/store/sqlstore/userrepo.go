package sqlstore

import (
	"go_blob_service/internal/app/model"

	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	store  *Store
	logger *logrus.Logger
}

func (r *UserRepository) Create(u *model.User) error {
	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) values ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)

	//r.logger.Info("USER CREATED")
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1", email,
	).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		return nil, err
	}

	return u, nil
}
