package utils

import (
	"time"
)

func GetISO8061(t time.Time) string {
	return t.Format(time.RFC3339)
}

func GetTimeFromString(tStr string) (time.Time, error) {
	return time.Parse(time.RFC3339, tStr)
}
