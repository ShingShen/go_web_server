package firebasemessaging

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func Connect(ctx context.Context) (*messaging.Client, error) {
	opt := option.WithCredentialsFile("config/env/firebase_key.json")

	firebaseApp, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("Failed to start firebase new app, err: %v", err)
	}

	firebaseMessaging, err := firebaseApp.Messaging(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to start firebase app messaging, err: %v", err)
	}

	return firebaseMessaging, nil
}

func SendMessaging(firebaseMessaging *messaging.Client, ctx context.Context, registrationToken string) {
	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Token: registrationToken,
	}

	response, err := firebaseMessaging.Send(ctx, message)
	if err != nil {
		fmt.Printf("Failed to send message, err: %v", err)
		// log.Fatalln(err)
	}
	fmt.Println("Successfully sent message:", response)
}

func SendNotification(firebaseMessaging *messaging.Client, ctx context.Context, registrationToken string) {
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Price drop",
			Body:  "5% off all electronics",
		},
		Token: registrationToken,
	}

	response, err := firebaseMessaging.Send(ctx, message)
	if err != nil {
		fmt.Printf("Failed to send message, err: %v", err)
		// log.Fatalln(err)
	}
	fmt.Println("Successfully sent message:", response)
}
