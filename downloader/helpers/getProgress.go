package helpers

import "fmt"

func GetProgress(downloaded int, total int) string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
	)

	var downloadedSize float32
	var totalSize float32

	var progress string

	switch {
	case total >= GB:
		downloadedSize = float32(downloaded) / float32(GB)
		totalSize = float32(total) / float32(GB)
		progress = fmt.Sprintf("%.1f/%.1f GB",downloadedSize,totalSize)
	case total >= MB:
		downloadedSize = float32(downloaded) / float32(MB)
		totalSize = float32(total) / float32(MB)
		progress = fmt.Sprintf("%.f/%.f MB",downloadedSize,totalSize)
	case total >= KB:
		downloadedSize = float32(downloaded) / float32(KB)
		totalSize = float32(total) / float32(KB)
		progress = fmt.Sprintf("%.f/%.f KB",downloadedSize,totalSize)
	default:
		downloadedSize = float32(downloaded)
		totalSize = float32(total)
		progress = fmt.Sprintf("%.f/%.f B",downloadedSize,totalSize)
	}

	return progress
}
