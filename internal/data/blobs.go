package data

import "database/sql"

type Blobs interface {
	Create(string, []byte, []byte, []byte) (string, error)
	Get(string) ([]byte, []byte, []byte, error)
	GetAll(string) (*sql.Rows, error)
	Delete(string) error
}
