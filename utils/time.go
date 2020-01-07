package utils

import "time"

func GetISO8061(t time.Time) string {
	return t.Format(time.RFC3339)
}
