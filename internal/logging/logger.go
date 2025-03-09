package logging

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
)

func InitLogger(logCategory string, logType string) *log.Logger {
	var logBuffer bytes.Buffer
	var logFile *os.File
	var logger *log.Logger
	var err error

	if logType == "console" {
		logFile = os.Stdout
	} else if logType == "file" {
		logFile, err = os.OpenFile("./logs/"+logType+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			fmt.Println("! Cannot open "+logType+" log for writing:", err)
			return nil
		}
	}
	logger = log.New(&logBuffer, ": ", log.Lshortfile)
	logger.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " - ")
	logger.SetOutput(logFile)

	return logger
}
