package logs

type LogStruct struct {
	Type      LogType
	Function  string
	Message   string
	UserId    int    // optional
	Endpoint  string // optional
	RequestId string // optional
}
