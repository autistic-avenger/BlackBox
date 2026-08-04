package helpers

import "fmt"

func GetSpeed(bytes int) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)
	var speed string
	switch {
	case bytes >= GB:
		speed = fmt.Sprintf("%.1f GB/s",float32(bytes)/GB)
	case bytes >= MB:
		speed = fmt.Sprintf("%.1f MB/s",float32(bytes)/MB)
	case bytes >= KB:
		speed = fmt.Sprintf("%.1f KB/s",float32(bytes)/KB)
	default:
		speed = fmt.Sprintf("%d B/s",bytes)
	}

	return speed
}