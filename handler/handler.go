package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func SaveItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	var p Payload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Printf("error: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	if err := SetItem(ctx, &p); err != nil {
		log.Printf("cannot set item %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
