package logs_test

import (
	"testing"

	"github.com/juanky-estevez/logs"
)

func init() {
	logs.LoadConfig("./logs", "UTC-5", "logs_test", "test")
}

func TestInfo(t *testing.T) {
	data := logs.LogStruct{
		Type:      logs.TypeInfo,
		Function:  "test/info",
		Message:   "Test message",
		UserId:    5,
		Endpoint:  "/test/info",
		RequestId: "X05",
	}

	logs.SaveLog(data)
}

func TestWarning(t *testing.T) {
	data := logs.LogStruct{
		Type:      logs.TypeWarning,
		Function:  "test/warning",
		Message:   "Test message",
		UserId:    5,
		Endpoint:  "/test/warning",
		RequestId: "X05",
	}

	logs.SaveLog(data)
}

func TestError(t *testing.T) {
	data := logs.LogStruct{
		Type:      logs.TypeError,
		Function:  "test/error",
		Message:   "Test message",
		UserId:    5,
		Endpoint:  "/test/error",
		RequestId: "X05",
	}

	logs.SaveLog(data)
}

func TestSuccess(t *testing.T) {
	data := logs.LogStruct{
		Type:      logs.TypeSuccess,
		Function:  "test/success",
		Message:   "Test message",
		UserId:    5,
		Endpoint:  "/test/success",
		RequestId: "X05",
	}

	logs.SaveLog(data)
}

func TestIncomplete(t *testing.T) {
	data := logs.LogStruct{
		Type:     logs.TypeInfo,
		Function: "test/incomplete",
		Message:  "Test message",
	}

	logs.SaveLog(data)
}
