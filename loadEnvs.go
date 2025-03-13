package logs

import "os"

var logFolder string
var timezone string
var projectName string
var environment string

func loadEnvs() {
	logsFolderAux, ok := os.LookupEnv("LOGS_FOLDER")
	if ok {
		logFolder = logsFolderAux
	} else {
		logFolder = "/logs"
	}

	timezoneAux, ok := os.LookupEnv("TIMEZONE")
	if ok {
		timezone = timezoneAux
	} else {
		timezone = "UTC-0"
	}

	projectNameAux, ok := os.LookupEnv("PROJECT_NAME")
	if ok {
		projectName = projectNameAux
	} else {
		projectName = "Default"
	}

	environmentAux, ok := os.LookupEnv("ENVIRONMENT")
	if ok {
		environment = environmentAux
	} else {
		environment = "no-env"
	}
}
