package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func Connect() *firebase.App {
	opt := option.WithCredentialsFile("config/env/firebase_key.json")

	firebaseApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	return firebaseApp
}
