package intervals

import "time"

type Period int

const(
	Unknown Period = iota
	FiveYears
	Year
	Month
	Week
	Day
)

func InSameInterval(t1 time.Time, t2 time.Time, p Period) bool {
	return FirstDay(t1, p) == FirstDay(t2, p)
}

func FirstDay(t time.Time, p Period) time.Time {
	var d day
	switch p {
	case Day:
		d = midnight(t)
	case Week:
		d = firstOfWeek(t)
	case Month:
		d = firstOfMonth(t)
	case Year:
		d = firstOfYear(t)
	case FiveYears:
		d = firstOfFiveYears(t)
	default:
		return time.Unix(0, 0)
	}
	return dayToTime(d)
}

func LastDay(t time.Time, p Period) time.Time {
	var d day
	switch p {
	case Day:
		d = midnight(t)
	case Week:
		d = lastOfWeek(t)
	case Month:
		d = lastOfMonth(t)
	case Year:
		d = lastOfYear(t)
	case FiveYears:
		d = lastOfFiveYears(t)
	default:
		return time.Unix(1<<63-1, 0)
	}
	return dayToTime(d)
}
