package main

import (
	"fmt"
	"log"
	"net/http"

	"webserver/internal/handlers"
	"webserver/internal/storage"
)

func main() {
	db, err := storage.NewSQLiteDB("./system_info.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/beacon", handlers.BeaconHandler(db))
	http.HandleFunc("/fetch", handlers.FetchHandler(db))

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
