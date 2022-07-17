package store

import "go_blob_service/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}

type BlobRepository interface {
	Create(*model.BlobMainData) error
	DeleteBlob(int) error
	GetBlobByID(int) (*model.BlobMainData, error)
	GetAllBlobs(int) ([]*model.BlobMainData, error)
	GenerateBlob(string, string, string, string, string, string) (*model.BlobMainData, error)
}
