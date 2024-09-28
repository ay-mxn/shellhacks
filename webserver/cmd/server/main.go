package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"webserver/internal/handlers"
	"webserver/internal/storage"
)

func main() {
	// Set up logging
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println("Starting server initialization...")

	// Initialize database
	log.Println("Connecting to SQLite database...")
	db, err := storage.NewSQLiteDB("./system_info.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Println("Successfully connected to SQLite database")

	// Set up HTTP server
	log.Println("Setting up HTTP routes...")
	http.HandleFunc("/beacon", logRequest(handlers.BeaconHandler(db)))
	http.HandleFunc("/fetch", logRequest(handlers.FetchHandler(db)))

	// Start the server
	addr := ":8080"
	log.Printf("Server is starting on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// logRequest is a middleware that logs incoming requests
func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Printf("Received %s request for %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s request for %s in %v", r.Method, r.URL.Path, time.Since(startTime))
	}
}
