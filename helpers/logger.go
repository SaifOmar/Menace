package helpers

import (
	"fmt"
	"os"
	"time"
)

type logger interface {
	Log(level, message string)
	Error(mesage string)
	Info(mesage string)
	GetLogs()
}

type TournamentLogger struct {
	logs      []string
	logToFile bool
	fileName  string
}

func NewTournamentLogger(logToFile bool, fileName string) *TournamentLogger {
	l := &TournamentLogger{
		logs:      []string{},
		logToFile: logToFile,
		fileName:  fileName,
	}
	return l
}

func (l *TournamentLogger) Log(level, message string) {
	timeStamp := time.Now().Local().Format(time.ANSIC)
	logEntry := fmt.Sprintf("[%s],[%s] : %s", timeStamp, level, message)
	l.logs = append(l.logs, logEntry)
	if l.logToFile {
		l.writeToFile(logEntry)
	}
}

func (l *TournamentLogger) Info(message string) {
	l.Log("INFO", message)
}

func (l *TournamentLogger) Error(message string) {
	l.Log("ERROR", message)
}

func (l *TournamentLogger) GetLogs() {
	fmt.Printf("%s", l.logs)
}

func (l *TournamentLogger) writeToFile(logEntry string) {
	file, err := os.OpenFile(l.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		l.Error(fmt.Sprintf("Error opening this file: %s", err))
		return
	}
	defer file.Close()
	_, err = file.WriteString(logEntry + "\n")
	if err != nil {
		l.Error(fmt.Sprintf("Error writing to this file: %s", err))
	}
}
