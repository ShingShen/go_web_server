package cloudstorage

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func Connect(ctx context.Context) (*storage.Client, error) {
	opt := option.WithCredentialsFile("config/env/firebase_key.json")

	cloudStorage, err := storage.NewClient(ctx, opt)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	return cloudStorage, nil
}
