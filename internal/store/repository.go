package store

import (
	"testService/internal/model"
	"testService/resources"
)

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}

type BlobRepository interface {
	Create(*resources.Blob) error
	DeleteBlob(string) error
	GetBlobByID(string) (*resources.BlobResponse, error)
	GetAllBlobs(string) (*resources.BlobListResponse, error)
	GenerateBlob(string, string, string, string,
		string, string, string, string, string) (*resources.Blob, error)
}
