package data

import "gitlab.com/distributed_lab/kit/pgdb"

type Blobs interface {
	Create(string, *Blob) (string, error)
	Get(string) (*Blob, error)
	GetAll() ([]Blob, error)
	Delete(string) error
	FilterByID(string) Blobs
	Page(pgdb.OffsetPageParams) Blobs
}

type Blob struct {
	ID            string `db:"id"`
	Key           []byte `db:"key"`
	Attributes    []byte `db:"attributes"`
	Relationships []byte `db:"relationships"`
}
