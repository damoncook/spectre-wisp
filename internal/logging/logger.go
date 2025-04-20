package logging

import (
	"bytes"
	"fmt"
	"log"
	"os"
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
	logger = log.New(&logBuffer, logCategory+": ", log.LstdFlags)
	//logger.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " - ")
	logger.SetOutput(logFile)

	return logger
}

func PrintErr(log *log.Logger, err error) {
	log.Printf("%s%s%s%s\n",
		Red,
		"ERROR: ",
		err,
		Reset)
}

func PrintGreen(log *log.Logger, msg string) {
	log.Printf("%s%s%s\n",
		Green,
		msg,
		Reset)
}
