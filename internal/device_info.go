package internal

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

type DeviceInfo struct {
	OS          string    `json:"os"`
	Hostname    string    `json:"hostname"`
	IPAddress   string    `json:"ip_address"`
	MacAddress  string    `json:"mac_address"`
	CPUCores    int       `json:"cpu_cores"`
	TotalMemory uint64    `json:"total_memory"`
	Timestamp   time.Time `json:"timestamp"`
	AccessType  string    `json:"access_type"`
}

func CollectAndSendDeviceInfo() error {
	accessType := determineAccessType()
	info := collectDeviceInfo(accessType)
	return sendDeviceInfo(info)
}

func collectDeviceInfo(accessType string) DeviceInfo {
	hostname, _ := os.Hostname()
	ipAddress := getIPAddress()
	macAddress := getMACAddress()

	return DeviceInfo{
		OS:          runtime.GOOS,
		Hostname:    hostname,
		IPAddress:   ipAddress,
		MacAddress:  macAddress,
		CPUCores:    runtime.NumCPU(),
		TotalMemory: getTotalMemory(),
		Timestamp:   time.Now(),
		AccessType:  accessType,
	}
}

func getIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func getMACAddress() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
			addrs, err := iface.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return iface.HardwareAddr.String()
					}
				}
			}
		}
	}
	return ""
}

func getTotalMemory() uint64 {
	v, err := mem.VirtualMemory()
	if err != nil {
		return 0
	}
	return v.Total
}

func sendDeviceInfo(info DeviceInfo) error {
	jsonData, err := json.Marshal(info)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8080/beacon", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func determineAccessType() string {
	if os.Getenv("SSH_CLIENT") != "" || os.Getenv("SSH_TTY") != "" {
		return "ssh"
	}

	proc, err := process.NewProcess(int32(os.Getppid()))
	if err == nil {
		name, err := proc.Name()
		if err == nil && strings.Contains(strings.ToLower(name), "curl") {
			return "curl"
		}
	}

	if isAtty() {
		return "terminal"
	}

	return "unknown"
}

func isAtty() bool {
	fileInfo, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}
