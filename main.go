package logs

import "log"

func SaveLog(data LogStruct) {
	shortDate, longDate := getDates()

	err := saveInFile(data, shortDate, longDate)
	if err != nil {
		log.Println("Error saving log in file:", err.Error())
	}

	showInTerminal(data, longDate)
}
