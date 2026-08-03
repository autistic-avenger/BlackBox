package helpers

import (
	"fmt"
	"time"
)

func CalculateSize(bytes int) string {
	const (
		KB = 1024
		MB = 1024 * 1024
		GB = 1024 * 1024 * 1024
	)

	size := float64(bytes)
	switch {
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", size/GB)
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", size/MB)
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", size/KB)
	default:
		return fmt.Sprintf("%d Bs", bytes)
	}
}

func CalculateTime(seconds int) string{
	return (time.Duration(seconds) * time.Second).String()	
}


func Truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	if max <= 3 {
		return s[:max]
	}
	return s[:max-3] + "..."
}