package dates

import "time"

const (
	dateLayout = "2006-02-01T15:44:05Z"
)

func GetCurrentDateTime() time.Time {
	return time.Now().UTC()
}

func GetCurrentDateTimeString() string {
	return GetCurrentDateTime().Format(dateLayout)
}
