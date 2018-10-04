package datetime

import (
	"./parser"
	"time"
)

// This
func GetDatetime(dt string, location *time.Location) (result time.Time, err error) {
	dateParams, eliminatedDt, err := parser.GetDate(dt)
	timeParams, _, err := parser.GetTime(eliminatedDt)

	result = time.Date(
		dateParams.Year,
		time.Month(dateParams.Month),
		dateParams.Day,
		timeParams.Hour,
		timeParams.Min,
		timeParams.Sec,
		timeParams.Nsec,
		location)

	return
}
