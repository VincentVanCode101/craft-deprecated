package utils

import (
	"os"
	"strings"
)

func IsRunningInDocker() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	file, err := os.Open("/proc/1/cgroup")
	if err != nil {
		return false
	}
	defer file.Close()

	buf := make([]byte, 4096)
	n, _ := file.Read(buf)
	return strings.Contains(string(buf[:n]), "docker")
}
