package intervals

import (
	"testing"
	"time"
)

var loc = time.UTC
var date = time.Date(2020, time.June, 24, 12, 12, 12, 12, loc) // Wed

var tableFirstDay = []struct {
	d      time.Time
	p      Period
	expect time.Time
}{
	{date, Day, time.Date(2020, time.June, 24, 0, 0, 0, 0, loc)},
	{date, Week, time.Date(2020, time.June, 22, 0, 0, 0, 0, loc)},
	{date, Month, time.Date(2020, time.June, 1, 0, 0, 0, 0, loc)},
	{date, Year, time.Date(2020, time.January, 1, 0, 0, 0, 0, loc)},
	{date, FiveYears, time.Date(2020, time.January, 1, 0, 0, 0, 0, loc)},
}

func TestFirstDay(t *testing.T) {
	for _, r := range tableFirstDay {
		got := FirstDay(r.d, r.p)
		if got != r.expect {
			t.Errorf("FirstDay(%v, %d) = %T", r.d, r.p, got)
		}
	}
}

var tableLastDay = []struct {
	d      time.Time
	p      Period
	expect time.Time
}{
	{date, Day, time.Date(2020, time.June, 24, 0, 0, 0, 0, loc)},
	{date, Week, time.Date(2020, time.June, 28, 0, 0, 0, 0, loc)},
	{date, Month, time.Date(2020, time.June, 30, 0, 0, 0, 0, loc)},
	{date, Year, time.Date(2020, time.December, 31, 0, 0, 0, 0, loc)},
	{date, FiveYears, time.Date(2024, time.December, 31, 0, 0, 0, 0, loc)},
}

func TestLastDay(t *testing.T) {
	for _, r := range tableLastDay {
		got := LastDay(r.d, r.p)
		if got != r.expect {
			t.Errorf("LastDay(%v, %s) = %v", r.d, r.p, got)
		}
	}
}
