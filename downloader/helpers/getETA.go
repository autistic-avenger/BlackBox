package helpers

import "fmt"

func GetETA(time int) string {
	const (
		MIN   = 60
		HOURS = MIN * 60
		DAY   = 24 * HOURS
		WEEK  = DAY * 7
	)

	switch {
	case time >= WEEK:
		week := time / WEEK
		day := time % WEEK
		if day != 0 {
			return fmt.Sprintf("%dw%dd",week,day/DAY)
		}
		return fmt.Sprintf("%dw",week)
	case time >= DAY:
		day := time / DAY
		hours := time % DAY
		if hours != 0 {
			return fmt.Sprintf("%dd%dh",day,hours/HOURS)
		}
		return fmt.Sprintf("%dd",day)
	case time >= HOURS:
		hours := time / HOURS
		mins := time % HOURS
		if mins != 0 {
			return fmt.Sprintf("%dh%dm",hours,mins/MIN)
		}
		return fmt.Sprintf("%dh",hours)

	case time >= MIN:
		mins := time / MIN
		sec := time % MIN
		if sec != 0 {
			return fmt.Sprintf("%dm%ds",mins,sec)
		}
		return fmt.Sprintf("%dm",mins)
	default:
		return fmt.Sprintf("%ds",time)
	}

}