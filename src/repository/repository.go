package repository

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/henzai/ranwei_functions/src/model"
)

const (
	COLLECTION_ITEM = "item"
)

func SetItem(ctx context.Context, p *model.Payload) error {
	pid := os.Getenv("PROJECT_ID")

	client, err := firestore.NewClient(ctx, pid)
	if err != nil {
		return err
	}

	i := model.PayloadToItem(p)
	_, err = client.Collection(COLLECTION_ITEM).NewDoc().Set(ctx, i)
	if err != nil {
		return err
	}
	return nil
}
