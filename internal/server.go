package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	db *SQLiteDB
}

func NewServer() (*Server, error) {
	log.Println("Initializing server...")
	db, err := NewSQLiteDB("./system_info.db")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %v", err)
	}
	log.Println("Database initialized successfully")
	return &Server{db: db}, nil
}

func (s *Server) Start() {
	log.Println("Setting up HTTP routes...")
	http.HandleFunc("/beacon", s.logRequest(s.BeaconHandler))
	http.HandleFunc("/fetch", s.logRequest(s.FetchHandler))

	addr := ":8080"
	log.Printf("Server is starting on http://localhost%s", addr)
	go func() {
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
}

func (s *Server) BeaconHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var info DeviceInfo
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	info.LastBeaconTime = time.Now()

	err = s.db.SaveSystemInfo(info)
	if err != nil {
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	log.Printf("Received and stored beacon from device: %s", info.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data received and stored successfully"))
}

func (s *Server) FetchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	info, err := s.db.GetSystemInfo(id)
	if err != nil {
		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
		return
	}

	if info == nil {
		http.Error(w, "ID not found", http.StatusNotFound)
		return
	}

	log.Printf("Fetched data for device: %s", id)
	json.NewEncoder(w).Encode(info)
}

func (s *Server) logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Printf("Received %s request for %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s request for %s in %v", r.Method, r.URL.Path, time.Since(startTime))
	}
}

type SQLiteDB struct {
	*sql.DB
}

func NewSQLiteDB(dbPath string) (*SQLiteDB, error) {
	log.Printf("Opening SQLite database at %s", dbPath)
	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	if err := createTable(conn); err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	log.Println("SQLite database initialized successfully")
	return &SQLiteDB{DB: conn}, nil
}

func createTable(conn *sql.DB) error {
	log.Println("Creating system_info table if it doesn't exist")
	_, err := conn.Exec(`
		CREATE TABLE IF NOT EXISTS system_info (
			id TEXT PRIMARY KEY,
			username TEXT,
			os TEXT,
			ram_total INTEGER,
			cpu_cores INTEGER,
			last_beacon_time DATETIME
		)
	`)
	return err
}

func (db *SQLiteDB) SaveSystemInfo(info DeviceInfo) error {
	log.Printf("Saving system info for ID: %s", info.ID)
	_, err := db.Exec(`INSERT OR REPLACE INTO system_info 
    (id, username, os, ram_total, cpu_cores, last_beacon_time) 
    VALUES (?, ?, ?, ?, ?, ?)`,
		info.ID, info.Username, info.OS, info.RAMTotal, info.CPUCores, info.LastBeaconTime)

	if err != nil {
		log.Printf("Error saving system info: %v", err)
	}
	return err
}

func (db *SQLiteDB) GetSystemInfo(id string) (*DeviceInfo, error) {
	log.Printf("Fetching system info for ID: %s", id)
	var info DeviceInfo
	err := db.QueryRow(`
		SELECT id, username, os, ram_total, cpu_cores, last_beacon_time 
		FROM system_info WHERE id = ?
	`, id).Scan(&info.ID, &info.Username, &info.OS, &info.RAMTotal, &info.CPUCores, &info.LastBeaconTime)

	if err == sql.ErrNoRows {
		log.Printf("No system info found for ID: %s", id)
		return nil, nil
	} else if err != nil {
		log.Printf("Error fetching system info: %v", err)
		return nil, err
	}

	log.Printf("Successfully fetched system info for ID: %s", id)
	return &info, nil
}
