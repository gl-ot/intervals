package intervals

import "time"

func firstOfYear(t time.Time) day {
	y, _, _ := t.Date()
	return day{y, 1, 1, t.Location()}
}

func lastOfYear(t time.Time) day {
	return timeToDay(dayToTime(firstOfYear(t)).AddDate(1, 0, -1))
}
