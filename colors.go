package logs

var colors = map[LogType]string{
	"error":   "\x1b[31m",
	"info":    "\x1b[34m",
	"success": "\x1b[32m",
	"warning": "\x1b[33m",
}
