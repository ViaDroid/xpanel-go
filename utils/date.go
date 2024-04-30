package utils

import "time"

type MyTime time.Time

// FirstDayOfMonth returns the first day of the current month at 00:00:00.
func FirstDayOfMonth() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

// GetTodayZeroTime returns today's date at 00:00:00.
func GetTodayZeroTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// GetTimeBeforeDays returns a time that is `days` days before the given `specificTime`.
func GetTimeBeforeDays(specificTime time.Time, days int) time.Time {
	return specificTime.AddDate(0, 0, -days) // Subtract days from the date part
}


// Get today's 1:00 time