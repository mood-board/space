package common

import (
	"strconv"
	"time"
)

var daysInMonth = map[time.Month]int{
	1:  31,
	2:  28,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

func ParseTimestamp(timestamp string) (time.Time, error) {
	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(i, 0), nil
}

func IsLeap() bool {
	year := time.Now().UTC().Year()
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func PreviousMonth(t time.Time) time.Time {
	// Since time.AddDate normalizes its result (e.g. adding a month to 31 Oct yields 1 Dec), it's not used here.
	var daysToSub int
	prevMonth := (t.Month()+10)%12 + 1
	if IsLeap() && prevMonth == 2 {
		daysToSub = 29
	} else {
		daysToSub = daysInMonth[prevMonth]
	}
	if t.Day() > daysToSub {
		daysToSub = t.Day()
	}
	return time.Date(t.Year(), t.Month(), t.Day()-daysToSub, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

func BeginningOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
}

func EndOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, -1, time.UTC)
}
