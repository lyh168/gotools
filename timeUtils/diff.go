package timeUtils

import "time"

func Differ(x, y time.Time) time.Duration {
	if x.Before(y) {
		return y.Sub(x)
	}
	return x.Sub(y)
}

func DifferWithDays(x, y time.Time, days int64) bool {
	duration := Differ(x, y)
	return duration.Nanoseconds() >= time.Hour.Nanoseconds()*24*days
}
