package intervals

import "time"

type day struct {
	year 	int
	month	time.Month
	day 	int
	loc 	*time.Location
}
