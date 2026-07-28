package helpers

import (
	"fmt"
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
	const (
		MIN = 60
		HOURS = MIN *60
		DAY = 24*HOURS
	)


	switch {
	case seconds >= DAY:
		return fmt.Sprintf("%d Days",seconds/DAY)
	case seconds >= HOURS:
		return fmt.Sprintf("%d Hours",seconds/HOURS)
	case seconds >= MIN:
		return fmt.Sprintf("%d Mins",seconds/MIN)
	default:
		return fmt.Sprintf("%d Sec", seconds)
	}
	
}
