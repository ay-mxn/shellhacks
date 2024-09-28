package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"webserver/internal/models"
	"webserver/internal/storage"
)

func BeaconHandler(db *storage.SQLiteDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var info models.SystemInfo
		err := json.NewDecoder(r.Body).Decode(&info)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		info.LastBeaconTime = time.Now()

		err = db.SaveSystemInfo(info)
		if err != nil {
			http.Error(w, "Failed to save data", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Data received and stored successfully"))
	}
}
