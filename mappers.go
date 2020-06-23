package intervals

import "time"

func dayToTime(d day) time.Time {
	return time.Date(d.year, d.month, d.day, 0, 0, 0, 0, d.loc)
}

func timeToDay(t time.Time) day {
	y, m, d := t.Date()
	return day{y, m, d, t.Location()}
}