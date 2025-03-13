package logs

func conditionalLogField(prefix, field string) string {
	if field == "" || field == "0" {
		return ""
	}

	return prefix + field + " - "
}
