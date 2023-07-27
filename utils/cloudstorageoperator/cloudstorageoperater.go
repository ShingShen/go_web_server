package cloudstorageoperator

import "cloud.google.com/go/storage"

type CloudStorage struct {
	CloudStorage *storage.Client
}

func (c CloudStorage) Bucket(name string) IBucketHandle {
	return c.CloudStorage.Bucket(name)
}
