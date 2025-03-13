package logs

func init() {
	loadEnvs()
	createDir()
	defineOffset()
}
