package handlers

import (
	"encoding/json"
	"net/http"
	"webserver/internal/storage"
)

func FetchHandler(db *storage.SQLiteDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing ID parameter", http.StatusBadRequest)
			return
		}

		info, err := db.GetSystemInfo(id)
		if err != nil {
			http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
			return
		}

		if info == nil {
			http.Error(w, "ID not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(info)
	}
}
