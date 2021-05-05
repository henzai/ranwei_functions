package functions

import (
	"net/http"

	"github.com/henzai/ranwei_functions/handler"
)

// SaveItem save attachment as item into firestore.
func SaveItem(w http.ResponseWriter, r *http.Request) {
	handler.SaveItem(w, r)
}
