package logs

import (
	"fmt"
)

func showInTerminal(data LogStruct, longDate string) {
	// Build message
	message := fmt.Sprintf(
		"\x1b[37m[%s%s\x1b[37m] \x1b[90m%s %s (%s) \x1b[37m%s%s%smessage: %s\x1b[0m",
		// "\x1b[37m[%s\x1b[0m%s\x1b[37m] \x1b[90m%s %s (%s)%s%s%s - message: \x1b[37m%s\x1b[0m",
		colors[data.Type],
		data.Type,
		environment,
		longDate,
		data.Function,
		conditionalLogField("endpoint: ", data.Endpoint),
		conditionalLogField("requestId: ", data.RequestId),
		conditionalLogField("userId: ", fmt.Sprintf("%d", data.UserId)),
		data.Message,
	)

	// Show in terminal
	fmt.Println(message)
}
