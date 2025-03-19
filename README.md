# logs

Installation:

```sh
go get github.com/juanky-estevez/logs
```

Configure example

```go
logsFolder := "./logs"
timezone := "UTC-5"
projectName := "logs-test"
environment := "test"

logs.LoadConfig(logsFolder, timezone, projectName, environment)
```

Default configuration values:

- logsFolder: "./logs"
- timezone: "UTC-0"
- projectName: "no_name"
- environment: "no_env"

Logs types:

- Info: logs.TypeInfo
- Warning: logs.TypeWarning
- Error: logs.TypeError
- Success: logs.TypeSuccess

How to use:

Using all options

```go
data := logs.LogStruct{
  Type:      logs.TypeInfo,
  Function:  "test_function",
  Message:   "Test message",
  UserId:    5,
  Endpoint:  "/test/info",
  RequestId: "X05",
}

logs.SaveLog(data)
```

Using only required options

```go
data := logs.LogStruct{
		Type:     logs.TypeInfo,
		Function: "test/incomplete",
		Message:  "Test message",
	}

	logs.SaveLog(data)
```

# Versions

## v1.0.0

First release
