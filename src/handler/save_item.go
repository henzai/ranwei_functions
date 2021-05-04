package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/henzai/ranwei_functions/src/model"
	"github.com/henzai/ranwei_functions/src/repository"
)

func SaveItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	var p model.Payload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Printf("error: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	if err := repository.SetItem(ctx, &p); err != nil {
		log.Printf("cannot set item %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
