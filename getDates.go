package logs

import "time"

// Returns: shortDate, longDate
func getDates() (string, string) {
	currentDateCentral := time.Now()
	currentDate := currentDateCentral.Add(time.Duration(offset) * time.Hour)

	// YYYYMMDD
	shortDate := currentDate.Format("20060102")

	// YYYY-MM-DD HH:MM:SS
	longDate := currentDate.Format("2006-01-02 15:04:05")

	return shortDate, longDate
}
