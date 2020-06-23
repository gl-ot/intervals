package intervals

import "time"

func firstOfWeek(t time.Time) day {
	y, m, d := t.Date()
	switch t.Weekday() {
	case time.Sunday:
		return day{y, m, d - 6, t.Location()}
	default:
		return day{y, m, d - int(t.Weekday()) + 1, t.Location()}
	}
}

func lastOfWeek(t time.Time) day {
	return timeToDay(dayToTime(firstOfWeek(t)).AddDate(0, 0, 7))
}
