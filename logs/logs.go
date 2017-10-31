package logs

import "github.com/cihub/seelog"

func LogsInit() {
	logger, err := seelog.LoggerFromConfigAsFile("./etc/seelog1.xml")
	if err != nil {
		seelog.Critical("err parsing config log file :", err)
		return
	}
	seelog.ReplaceLogger(logger)
	defer seelog.Flush()
}