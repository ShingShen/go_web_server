package firestoredb

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func Connect() (*firestore.Client, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("config/env/firebase_key.json")

	firebaseApp, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	firestoreDB, err := firebaseApp.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return firestoreDB, nil
}
