package helpers

import (
	"time"
)

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