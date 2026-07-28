package helpers

import "fmt"

func CalculateSize(bytes int) string {
	const (
		KB = 1024
		MB = 1024 * 1024
		GB = 1024 * 1024 * 1024
	)

	size := float64(bytes)
	switch {
	case bytes >= GB:
		return fmt.Sprintf("%.2f GBs", size/GB)
	case bytes >= MB:
		return fmt.Sprintf("%.2f MBs", size/MB)
	case bytes >= KB:
		return fmt.Sprintf("%.2f KBs", size/KB)
	default:
		return fmt.Sprintf("%d Bs", bytes)
	}
}