package sqlstore

import (
	"errors"
	"testService/internal/service/helpers"
	"testService/resources"

	"github.com/sirupsen/logrus"
)

const url string = "http://localhost:8081/integrations/testService/blob/"

type BlobRepository struct {
	store  *Store
	logger *logrus.Logger
}

func (r *BlobRepository) GenerateBlob(
	Type string,
	Title string,
	LinkFirst string,
	LinkLast string,
	LinkPrev string,
	LinkSelf string,
	LinkNext string,
	AuthorType string,
	AuthorID string) (*resources.Blob, error) {

	return &resources.Blob{
		Key: resources.Key{
			Type: resources.ResourceType(Type),
		},
		Attributes: resources.BlobAttributes{
			Title: Title,
		},
		Relationships: resources.BlobRelationships{
			Author: resources.Relation{
				Data: &resources.Key{
					Type: resources.ResourceType(AuthorType),
					ID:   AuthorID,
				},
				Links: &resources.Links{
					First: LinkFirst,
					Last:  LinkLast,
					Next:  LinkNext,
					Prev:  LinkPrev,
					Self:  LinkSelf,
				},
			},
		},
	}, nil
}

func (r *BlobRepository) DeleteBlob(id string) error {
	res, err := r.store.db.Exec("DELETE FROM blobs WHERE id=$1", id)
	aff_rows, _ := res.RowsAffected()
	if err != nil {
		return err
	}
	if aff_rows < 1 {
		return errors.New("THERE IS NO SUCH BLOB")
	}

	r.logger.Info("BLOB DELETED")
	return nil
}

func (r *BlobRepository) Create(blob *resources.Blob) error {
	attributes, err := helpers.JSON_stringify(blob.Attributes)
	if err != nil {
		r.logger.Error("ERROR JSON STRINGIFY")
		return err
	}
	relationships, err := helpers.JSON_stringify(blob.Relationships)
	if err != nil {
		r.logger.Error("ERROR JSON STRINGIFY")
		return err
	}
	key, err := helpers.JSON_stringify(blob.Key)
	if err != nil {
		r.logger.Error("ERROR JSON STRINGIFY")
		return err
	}

	err = r.store.db.QueryRow(
		"INSERT INTO blobs (user_id, key, attributes, relationships) values ($1, $2, $3, $4) RETURNING id",
		blob.Relationships.Author.Data.ID,
		key,
		attributes,
		relationships,
	).Scan(&blob.Key.ID)

	if err != nil {
		return err
	}

	r.logger.Info("BLOB CREATED")
	return nil
}

func (r *BlobRepository) GetBlobByID(id string) (*resources.BlobResponse, error) {
	blob := resources.Blob{}

	var key []byte
	var attributes []byte
	var relationships []byte

	if err := r.store.db.QueryRow(
		"SELECT key, attributes, relationships FROM blobs WHERE id = $1", id,
	).Scan(&key, &attributes, &relationships); err != nil {
		return nil, err
	}

	if err := helpers.JSON_parse(attributes, &blob.Attributes); err != nil {
		r.logger.Error("ERROR PARSING JSON")
		return nil, err
	}
	if err := helpers.JSON_parse(relationships, &blob.Relationships); err != nil {
		r.logger.Error("ERROR PARSING JSON")
		return nil, err
	}
	if err := helpers.JSON_parse(key, &blob.Key); err != nil {
		r.logger.Error("ERROR PARSING JSON")
		return nil, err
	}
	blob.Key.ID = id

	r.logger.Info("BLOB FOUND")

	return &resources.BlobResponse{
		Data:     blob,
		Included: resources.Included{},
	}, nil
}

func (r *BlobRepository) GetAllBlobs(id string) (*resources.BlobListResponse, error) {
	blobs := make([]resources.Blob, 0)
	links := &resources.Links{}

	rows, err := r.store.db.Query(
		"SELECT id, key, attributes, relationships FROM blobs WHERE user_id = $1", id,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	counter := 0
	for rows.Next() {

		var key []byte
		var attributes []byte
		var relationships []byte
		var id string

		blob := resources.Blob{}

		if err := rows.Scan(&id, &key, &attributes, &relationships); err != nil {
			r.logger.Error("ERROR ROWS PROCESSING")
			return nil, err
		}

		if err := helpers.JSON_parse(attributes, &blob.Attributes); err != nil {
			r.logger.Error("ERROR PARSING JSON")
			return nil, err
		}
		if err := helpers.JSON_parse(relationships, &blob.Relationships); err != nil {
			r.logger.Error("ERROR PARSING JSON")
			return nil, err
		}
		if err := helpers.JSON_parse(key, &blob.Key); err != nil {
			r.logger.Error("ERROR PARSING JSON")
			return nil, err
		}

		if counter == 0 {
			links.Self = url + id
			links.First = links.Self
			links.Prev = links.Self
		}

		if counter == 1 {
			links.Next = url + id
		}

		blob.Key.ID = id
		blobs = append(blobs, blob)
		counter++
	}

	links.Last = url + blobs[len(blobs)-1].Key.ID

	if len(blobs) < 1 {
		return nil, errors.New("NO BLOBS FOUND")
	}

	r.logger.Info("ALL BLOBS FOUND")
	return &resources.BlobListResponse{
		Data:     blobs,
		Included: resources.Included{},
		Links:    links,
	}, nil
}
