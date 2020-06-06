package log

import log "github.com/cihub/seelog"

func Debug(v ...interface{}) {
	logger, _ := log.LoggerFromConfigAsFile("conf/seelog.xml")
	defer log.Flush()
	log.ReplaceLogger(logger)
	log.Debug(v)
}
