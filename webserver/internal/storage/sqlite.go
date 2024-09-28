package storage

import (
	"database/sql"
	"log"
	"webserver/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
	*sql.DB
}

func NewSQLiteDB(dbPath string) (*SQLiteDB, error) {
	log.Printf("Opening SQLite database at %s", dbPath)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Printf("Failed to open database: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Printf("Failed to ping database: %v", err)
		return nil, err
	}

	if err := createTable(db); err != nil {
		log.Printf("Failed to create table: %v", err)
		return nil, err
	}

	log.Println("SQLite database initialized successfully")
	return &SQLiteDB{DB: db}, nil
}

func createTable(db *sql.DB) error {
	log.Println("Creating system_info table if it doesn't exist")
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS system_info (
			id TEXT PRIMARY KEY,
			username TEXT,
			os TEXT,
			ram_total INTEGER,
			cpu_cores INTEGER,
			file_count INTEGER,
			last_beacon_time DATETIME
		)
	`)
	return err
}

func (db *SQLiteDB) SaveSystemInfo(info models.SystemInfo) error {
	log.Printf("Saving system info for ID: %s", info.ID)
	_, err := db.Exec(`
		INSERT OR REPLACE INTO system_info 
		(id, username, os, ram_total, cpu_cores, file_count, last_beacon_time) 
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, info.ID, info.Username, info.OS, info.RAMTotal, info.CPUCores, info.FileCount, info.LastBeaconTime)
	if err != nil {
		log.Printf("Error saving system info: %v", err)
	}
	return err
}

func (db *SQLiteDB) GetSystemInfo(id string) (*models.SystemInfo, error) {
	log.Printf("Fetching system info for ID: %s", id)
	var info models.SystemInfo
	err := db.QueryRow(`
		SELECT id, username, os, ram_total, cpu_cores, file_count, last_beacon_time 
		FROM system_info WHERE id = ?
	`, id).Scan(&info.ID, &info.Username, &info.OS, &info.RAMTotal, &info.CPUCores, &info.FileCount, &info.LastBeaconTime)

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
