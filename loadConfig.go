package logs

var folderLogs string
var timezoneLogs string
var projectNameLogs string
var environmentLogs string

func LoadConfig(
	logsFolder string,
	timezone string,
	projectName string,
	environment string,
) {
	if logsFolder != "" {
		folderLogs = logsFolder
	} else {
		folderLogs = "./logs"
	}

	if timezone != "" {
		timezoneLogs = timezone
	} else {
		timezoneLogs = "UTC-0"
	}

	if projectName != "" {
		projectNameLogs = projectName
	} else {
		projectNameLogs = "no_name"
	}

	if environment != "" {
		environmentLogs = environment
	} else {
		environmentLogs = "no_env"
	}

	createDir()
	defineOffset()
}
