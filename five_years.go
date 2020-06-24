package intervals

import "time"

func firstOfFiveYears(t time.Time) day {
	y, _, _ := t.Date()
	diff := y % 10
	if diff > 5 {
		diff += 5
	}
	rightYear := y - diff
	return day{rightYear, 1, 1, t.Location()}
}

func lastOfFiveYears(t time.Time) day {
	return timeToDay(dayToTime(firstOfYear(t)).AddDate(5, 0, -1))
}