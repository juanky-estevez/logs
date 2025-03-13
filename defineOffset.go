package logs

import (
	"strconv"
	"strings"
)

var offset = 0

func defineOffset() {
	if timezone == "" || strings.Contains(timezone, "UTC") {
		offset = 0
	}

	offsetString := strings.ReplaceAll(timezone, "UTC", "")
	if offsetString == "" {
		offset = 0
	}

	offsetAux, err := strconv.Atoi(offsetString)
	if err != nil {
		offset = 0
	} else {
		offset = offsetAux
	}
}
