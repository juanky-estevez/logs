package logs

import (
	"fmt"
	"os"
	"strings"
)

func saveInFile(data LogStruct, shortDate string, longDate string) error {

	filename := fmt.Sprintf("%s_%s_%s.log", projectNameLogs, environmentLogs, shortDate)

	message := fmt.Sprintf(
		"[%s] %s %s (%s) %s%s%smessage: %s \n",
		data.Type,
		environmentLogs,
		longDate,
		data.Function,
		conditionalLogField("endpoint: ", data.Endpoint),
		conditionalLogField("requestId: ", data.RequestId),
		conditionalLogField("userId: ", fmt.Sprintf("%d", data.UserId)),
		data.Message,
	)

	// Create log file
	var filePath string
	if strings.HasSuffix(folderLogs, "/") {
		filePath = folderLogs + filename
	} else {
		filePath = folderLogs + "/" + filename
	}

	// Open file
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Save message
	_, err = file.WriteString(message)
	if err != nil {
		return err
	}

	return nil
}
