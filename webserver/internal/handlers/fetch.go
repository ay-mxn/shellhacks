package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"webserver/internal/storage"
)

func FetchHandler(db *storage.SQLiteDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Printf("Method not allowed: %s", r.Method)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Query().Get("id")
		if id == "" {
			log.Print("Missing ID parameter")
			http.Error(w, "Missing ID parameter", http.StatusBadRequest)
			return
		}

		log.Printf("Fetching data for ID: %s", id)

		info, err := db.GetSystemInfo(id)
		if err != nil {
			log.Printf("Failed to retrieve data for ID %s: %v", id, err)
			http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
			return
		}

		if info == nil {
			log.Printf("ID not found: %s", id)
			http.Error(w, "ID not found", http.StatusNotFound)
			return
		}

		log.Printf("Successfully retrieved data for ID: %s", id)
		json.NewEncoder(w).Encode(info)
	}
}
