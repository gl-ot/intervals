package intervals

import "time"

func midnight(t time.Time) day {
	y, m, d := t.Date()
	return day{y, m, d, t.Location()}
}
