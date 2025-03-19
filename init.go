package logs

import (
	"os"
)

var folderLogs string
var timezoneLogs string
var projectNameLogs string
var environmentLogs string

func init() {
	folderLogsEnv, ok := os.LookupEnv("LOGS_FOLDER")
	if !ok || folderLogsEnv == "" {
		folderLogs = "./logs"
	} else {
		folderLogs = folderLogsEnv
	}

	timezoneEnv, ok := os.LookupEnv("TIMEZONE")
	if !ok || timezoneEnv == "" {
		timezoneLogs = "UTC-0"
	} else {
		timezoneLogs = timezoneEnv
	}

	projectNameEnv, ok := os.LookupEnv("PROJECT_NAME")
	if !ok || projectNameEnv == "" {
		projectNameLogs = "no-name"
	} else {
		projectNameLogs = projectNameEnv
	}

	environmentEnv, ok := os.LookupEnv("ENVIRONMENT")
	if !ok || environmentEnv == "" {
		environmentLogs = "no-env"
	} else {
		environmentLogs = environmentEnv
	}

	createDir()
	defineOffset()
}
