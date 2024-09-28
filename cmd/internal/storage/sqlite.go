package storage

import (
	"database/sql"
	"webserver/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
	*sql.DB
}

func NewSQLiteDB(dbPath string) (*SQLiteDB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createTable(db); err != nil {
		return nil, err
	}

	return &SQLiteDB{DB: db}, nil
}

func createTable(db *sql.DB) error {
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
	_, err := db.Exec(`
		INSERT OR REPLACE INTO system_info 
		(id, username, os, ram_total, cpu_cores, file_count, last_beacon_time) 
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, info.ID, info.Username, info.OS, info.RAMTotal, info.CPUCores, info.FileCount, info.LastBeaconTime)
	return err
}

func (db *SQLiteDB) GetSystemInfo(id string) (*models.SystemInfo, error) {
	var info models.SystemInfo
	err := db.QueryRow(`
		SELECT id, username, os, ram_total, cpu_cores, file_count, last_beacon_time 
		FROM system_info WHERE id = ?
	`, id).Scan(&info.ID, &info.Username, &info.OS, &info.RAMTotal, &info.CPUCores, &info.FileCount, &info.LastBeaconTime)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &info, nil
}
