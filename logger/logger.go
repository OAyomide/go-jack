package logger

import (
	log "github.com/sirupsen/logrus"
)

type Logger struct {
	message string
	level   string
}

//a function to apply colorization to our log
func Colorise(level string) {
	l := log.WithFields(log.Fields{
		"event": "Max",
	})

	l.Info("Hello there!")
}
