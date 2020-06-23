package intervals

import "time"

func firstOfMonth(t time.Time) day {
	y, m, _ := t.Date()
	return day{y, m, 1, t.Location()}
}

func lastOfMonth(t time.Time) day {
	return timeToDay(dayToTime(firstOfMonth(t)).AddDate(0, 0, -1))
}
