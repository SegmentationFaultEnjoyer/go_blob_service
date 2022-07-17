package store

type Store interface {
	User() UserRepository
	Blob() BlobRepository
}
