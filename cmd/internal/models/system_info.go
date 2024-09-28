package models

import (
	"time"
)

type SystemInfo struct {
	ID             string    `json:"id"`
	Username       string    `json:"username"`
	OS             string    `json:"os"`
	RAMTotal       int       `json:"ram_total"`
	CPUCores       int       `json:"cpu_cores"`
	FileCount      int       `json:"file_count"`
	LastBeaconTime time.Time `json:"last_beacon_time"`
}
