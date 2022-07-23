package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	blobsTable = "blobs"
)

var (
	blobsSelect = squirrel.
		Select("key", "attributes", "relationships").
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

func (q *Blobs) Create(user_id string, atr []byte, rel []byte, key []byte) (string, error) {
	query := squirrel.Insert(blobsTable).
		Columns("user_id", "key", "attributes", "relationships").
		Values(user_id, key, atr, rel).
		Suffix("RETURNING \"id\"").
		RunWith(q.db.RawDB()).
		PlaceholderFormat(squirrel.Dollar)

	var id string

	if err := query.QueryRow().Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (q *Blobs) Get(id string) ([]byte, []byte, []byte, error) {
	var key []byte
	var attributes []byte
	var relationships []byte

	query := q.stmt.Where("id = ?", id).
		RunWith(q.db.RawDB()).
		PlaceholderFormat(squirrel.Dollar)

	if err := query.QueryRow().Scan(&key, &attributes, &relationships); err != nil {
		return nil, nil, nil, err
	}

	return key, attributes, relationships, nil
}

func (q *Blobs) GetAll(id string) (*sql.Rows, error) {
	var query squirrel.SelectBuilder

	if id == "" {
		query = squirrel.
			Select("id, key, attributes, relationships").
			From(blobsTable).
			RunWith(q.db.RawDB())
	} else {
		query = squirrel.
			Select("id, key, attributes, relationships").
			From(blobsTable).
			Where("user_id = ?", id).
			RunWith(q.db.RawDB()).
			PlaceholderFormat(squirrel.Dollar)
	}

	rows, err := query.Query()

	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (q *Blobs) Delete(id string) error {
	if _, _, _, err := q.Get(id); err != nil {
		return err
	}

	stmt := squirrel.
		Delete(blobsTable).
		Where("id = ?", id)

	if err := q.db.Exec(stmt); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
