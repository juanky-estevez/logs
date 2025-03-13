package logs

type LogType string

const (
	TypeInfo    LogType = "info"
	TypeWarning LogType = "warning"
	TypeError   LogType = "error"
	TypeSuccess LogType = "success"
)
