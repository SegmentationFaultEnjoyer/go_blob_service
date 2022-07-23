package helpers

import "testService/resources"

func GenerateBlob(
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
