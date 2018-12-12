package logger

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	message string
	level   string
}

//a function to apply colorization to our log
func colorise(level string) {
	l := log.WithFields(log.Fields{
		"event": "Max",
	})

	l.Info("Hello!")
	l.Warn("This is a warning")
}

//a function that prints the data we pass to
func Print(level string, message string, rawdata string) {
	l := log.WithFields(log.Fields{
		"[ACTION]": time.Now().Format("2010-01-02 3:4:5 pm"),
	})

	basicEntry := Logger{
		message: message,
		level:   level,
	}

	lg, _ := json.Marshal(basicEntry)
	logData := string(lg)
	switch level {
	case "info":
		l.Info(logData)

	case "warning":
		l.Warning(logData)

	case "debug":
		l.Debug(logData)

	case "error":
		l.Error(logData)
	}
}

//creates the logfile folders 📁
func CreateLogFolder(fileName string) {
	err := os.Mkdir(fileName, 777)

	if err != nil {
		Print("error", "Folder creating error", err.Error())
	}
}

func CreateLogFile(folderName string, fileName string) {
	_, err := os.Create(fileName)

	if err != nil {
		Print("error", "File not created", err.Error())
	}
}

func WriteToLogFile(fileName string, data string) {
	err := ioutil.WriteFile(fileName, []byte(data), 0666)

	if err != nil {
		Print("error", "Writing error", err.Error())
	}
}
