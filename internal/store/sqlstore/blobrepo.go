package sqlstore

import (
	"errors"
	"strconv"
	"testService/internal/model"
	"testService/internal/service/helpers"

	"github.com/sirupsen/logrus"
)

type BlobRepository struct {
	store  *Store
	logger *logrus.Logger
}

func (r *BlobRepository) GenerateBlob(Type string, Title string, LinkSelf string,
	LinkRelated string, AuthorType string, AuthorID string) (*model.BlobMainData, error) {
	id, err := strconv.Atoi(AuthorID)
	if err != nil {
		return nil, err
	}

	return &model.BlobMainData{
		Type: Type,
		Attributes: &model.BlobAtr{
			Title: Title,
		},
		Relationships: &model.BlobRelationships{
			Author: &model.BlobAuthor{
				Links: &model.BlobLinks{
					Self:    LinkSelf,
					Related: LinkRelated,
				},
				Data: &model.BlobData{
					Type: AuthorType,
					ID:   id,
				},
			},
		},
	}, nil
}

func (r *BlobRepository) DeleteBlob(id int) error {
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

func (r *BlobRepository) Create(blob *model.BlobMainData) error {
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

	err = r.store.db.QueryRow(
		"INSERT INTO blobs (user_id, type, attributes, relationships) values ($1, $2, $3, $4) RETURNING id",
		blob.Relationships.Author.Data.ID,
		blob.Type,
		attributes,
		relationships,
	).Scan(&blob.ID)

	if err != nil {
		return err
	}

	r.logger.Info("BLOB CREATED")
	return nil
}

func (r *BlobRepository) GetBlobByID(id int) (*model.BlobMainData, error) {
	blob := &model.BlobMainData{}

	var attributes []byte
	var relationships []byte
	if err := r.store.db.QueryRow(
		"SELECT id, type, attributes, relationships FROM blobs WHERE id = $1", id,
	).Scan(&blob.ID, &blob.Type, &attributes, &relationships); err != nil {
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

	r.logger.Info("BLOB FOUND")

	return blob, nil
}

func (r *BlobRepository) GetAllBlobs(id int) ([]*model.BlobMainData, error) {
	blobs := make([]*model.BlobMainData, 0)
	rows, err := r.store.db.Query(
		"SELECT id, type, attributes, relationships FROM blobs WHERE user_id = $1", id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var attributes []byte
		var relationships []byte

		blob := &model.BlobMainData{}
		if err := rows.Scan(&blob.ID, &blob.Type, &attributes, &relationships); err != nil {
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

		blobs = append(blobs, blob)
	}

	if len(blobs) < 1 {
		return nil, errors.New("NO BLOBS FOUND")
	}

	r.logger.Info("ALL BLOBS FOUND")
	return blobs, nil
}
