package logger

import (
	log "github.com/jeanphorn/log4go"
	"queueup/libs/config"
)

type logger struct {
}

var ins *logger

func New() *logger {
	if ins == nil {
		log.LoadConfiguration(config.ConfigPath + "/" + "logger.json")
		ins = new(logger)
	}
	return ins
}

func (self *logger) Info(level string, message string) bool {
	log.LOGGER(level).Info(message)
	return true
}
