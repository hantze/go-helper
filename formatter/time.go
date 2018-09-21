package formatter

import (
	"fmt"
	"time"
)

// FirstDay ...
func FirstDay(year int, month int, timezone int) (time.Time, string) {
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.FixedZone("time", timezone*3600))
	startDate := fmt.Sprintf("%d-%02d-%02d", firstDay.Year(), firstDay.Month(), firstDay.Day())
	return firstDay, startDate
}

// LastDay ...
func LastDay(firstDay time.Time) (time.Time, string) {
	endDay := firstDay.AddDate(0, 1, 0).Add(-time.Nanosecond)
	endDate := fmt.Sprintf("%d-%02d-%02d", endDay.Year(), endDay.Month(), endDay.Day())
	return endDay, endDate
}

// LastTwoMonths ...
func LastTwoMonths(firstDay time.Time) (time.Time, string) {
	endDay := firstDay.AddDate(0, 2, 0).Add(-time.Nanosecond)
	endDate := fmt.Sprintf("%d-%02d-%02d", endDay.Year(), endDay.Month(), endDay.Day())
	return endDay, endDate
}

// Expired ...
func Expired(start, end time.Time) bool {
	return start.Before(end)
}

// GetNow ...
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowAsiaJakarta ...
func GetNowAsiaJakarta() time.Time {
	utc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	return time.Now().In(utc)
}
