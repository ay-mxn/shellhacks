package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type DeviceInfo struct {
	ID             string    `json:"id"`
	Username       string    `json:"username"`
	OS             string    `json:"os"`
	RAMTotal       int       `json:"ram_total"`
	CPUCores       int       `json:"cpu_cores"`
	LastBeaconTime time.Time `json:"last_beacon_time"`
}

func CollectAndSendDeviceInfo() error {
	log.Println("Collecting device information...")
	info, err := collectDeviceInfo()
	if err != nil {
		return fmt.Errorf("failed to collect device info: %v", err)
	}

	log.Println("Sending device information to server...")
	return sendDeviceInfo(info)
}

func collectDeviceInfo() (DeviceInfo, error) {
	username := os.Getenv("USER")
	if username == "" {
		username = os.Getenv("USERNAME")
	}

	v, err := mem.VirtualMemory()
	if err != nil {
		return DeviceInfo{}, fmt.Errorf("failed to get virtual memory info: %v", err)
	}

	h, err := host.Info()
	if err != nil {
		return DeviceInfo{}, fmt.Errorf("failed to get host info: %v", err)
	}

	log.Printf("Collected device info: ID=%s, Username=%s, OS=%s, RAM=%d MB, CPUs=%d",
		h.HostID, username, runtime.GOOS, v.Total/1024/1024, runtime.NumCPU())

	return DeviceInfo{
		ID:             h.HostID,
		Username:       username,
		OS:             runtime.GOOS,
		RAMTotal:       int(v.Total / 1024 / 1024), // Convert to MB
		CPUCores:       runtime.NumCPU(),
		LastBeaconTime: time.Now(),
	}, nil
}

func sendDeviceInfo(info DeviceInfo) error {
	jsonData, err := json.Marshal(info)
	if err != nil {
		return fmt.Errorf("failed to marshal device info: %v", err)
	}

	resp, err := http.Post("http://localhost:8080/beacon", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send device info: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned non-OK status: %s", resp.Status)
	}

	log.Println("Device information sent successfully")
	return nil
}
