package datetime

import (
	"time"
)

// This
func GetDatetime(dt string, location *time.Location) (result time.Time, err error) {
	dateParams, eliminatedDt, err := getDate(dt)
	timeParams, _, err := getTime(eliminatedDt)

	result = time.Date(
		dateParams.year,
		time.Month(dateParams.month),
		dateParams.day,
		timeParams.hour,
		timeParams.min,
		timeParams.sec,
		timeParams.nsec,
		location)

	return
}
