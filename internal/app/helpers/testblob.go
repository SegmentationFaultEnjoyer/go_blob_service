package helpers

import "go_blob_service/internal/app/model"

func GetTestBlob() *model.BlobMainData {
	return &model.BlobMainData{
		Type: "article",
		Attributes: &model.BlobAtr{
			Title: "test_title",
		},
		Relationships: &model.BlobRelationships{
			Author: &model.BlobAuthor{
				Links: &model.BlobLinks{
					Self:    "href",
					Related: "/related",
				},
				Data: &model.BlobData{
					Type: "people",
					ID:   1,
				},
			},
		},
	}
}
