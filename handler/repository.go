package handler

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

const (
	COLLECTION_ITEMS = "items"
)

var pid = os.Getenv("PROJECT_ID")

var client *firestore.Client

func init() {
	// err is pre-declared to avoid shadowing client.
	var err error

	client, err = firestore.NewClient(context.Background(), pid)
	if err != nil {
		log.Fatalf("firestore.NewClient: %v", err)
	}
}

func SetItem(ctx context.Context, p *Payload) error {
	i := PayloadToItem(p)
	_, err := client.Collection(COLLECTION_ITEMS).NewDoc().Set(ctx, i)
	if err != nil {
		return err
	}
	return nil
}
