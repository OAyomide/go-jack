package logger

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
)

//a function to apply colorization to our log
func colorise(level string) {
	l := log.WithFields(log.Fields{
		"event": "Max",
	})

	l.Info("Hello!")
	l.Warn("This is a warning")
}

//Print is a function that prints the data we pass to
func Print(level string, message string, rawdata ...string) {
	log.SetFormatter(&log.JSONFormatter{PrettyPrint: true, DisableTimestamp: true})

	l := log.WithFields(log.Fields{
		"[TIMESTAMP]": time.Now().Format("2010-01-02 3:4:5 pm"),
		"raw":         rawdata,
	})

	switch level {
	case "info":
		l.Info(message)

	case "warning":
		l.Warning(message)

	case "debug":
		l.Debug(message)

	case "error":
		l.Error(message)
	}
}

//CreateLogFolder creates the logfile folders 📁
func CreateLogFolder(fileName string) {
	err := os.Mkdir("logs"+"/"+fileName, 777)

	if err != nil {
		Print("error", "Folder creating error", err.Error())
	}
}

//CreateLogFile creates the logfiles
func CreateLogFile(folderName string, fileName string) {
	filePath := filepath.FromSlash("logs" + "/" + folderName + "/" + fileName)
	_, err := os.Create(filePath)

	if err != nil {
		Print("error", "File not created", err.Error())
	}
}

//WriteToLogFile writes to the created LogFile
func WriteToLogFile(fileName string, data string) {
	err := ioutil.WriteFile(fileName, []byte(data), 0666)

	if err != nil {
		Print("error", "Writing error", err.Error())
	}
}
