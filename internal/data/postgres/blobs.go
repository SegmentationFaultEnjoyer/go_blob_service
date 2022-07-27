package postgres

import (
	"errors"
	"testService/internal/data"

	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	blobsTable = "blobs"
)

var (
	blobsSelect = squirrel.
		Select("key", "attributes", "relationships", "id").
		From(blobsTable)
)

type Blobs struct {
	db   *pgdb.DB
	stmt squirrel.SelectBuilder
}

func New(db *pgdb.DB) *Blobs {
	return &Blobs{
		db:   db,
		stmt: blobsSelect,
	}
}

func (q *Blobs) Create(user_id string, blobData *data.Blob) (string, error) {
	query := squirrel.Insert(blobsTable).
		Columns("user_id", "key", "attributes", "relationships").
		Values(user_id, blobData.Key, blobData.Attributes, blobData.Relationships).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar)

	var id string

	if err := q.db.Get(&id, query); err != nil {
		return "", err
	}

	return id, nil
}

func (q *Blobs) Get(id string) (*data.Blob, error) {
	var result data.Blob

	query := q.stmt.
		Where("id = ?", id).
		PlaceholderFormat(squirrel.Dollar)

	if err := q.db.Get(&result, query); err != nil {
		return nil, err
	}

	return &result, nil
}

func (q *Blobs) FilterByID(id string) data.Blobs {
	q.stmt = q.stmt.
		Where("user_id = ?", id).
		PlaceholderFormat(squirrel.Dollar)
	return q
}

func (q *Blobs) Page(pageParams pgdb.OffsetPageParams) data.Blobs {
	q.stmt = pageParams.ApplyTo(q.stmt, "id")
	return q
}

func (q *Blobs) GetAll() ([]data.Blob, error) {
	var results []data.Blob

	if err := q.db.Select(&results, q.stmt); err != nil {
		return nil, err
	}

	if len(results) < 1 {
		return nil, errors.New("No blobs found")
	}

	q.stmt = blobsSelect
	return results, nil
}

func (q *Blobs) Delete(id string) error {
	if _, err := q.Get(id); err != nil {
		return err
	}

	stmt := squirrel.
		Delete(blobsTable).
		Where("id = ?", id)

	if err := q.db.Exec(stmt); err != nil {
		return err
	}

	return nil
}
