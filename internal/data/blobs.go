package data

type Blobs interface {
	Create(string, *Blob) (string, error)
	Get(string) (*Blob, error)
	GetAll() ([]Blob, error)
	Delete(string) error
	FilterByID(string) Blobs
}

type Blob struct {
	ID            string `db:"id"`
	Key           []byte `db:"key"`
	Attributes    []byte `db:"attributes"`
	Relationships []byte `db:"relationships"`
}
