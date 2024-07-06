package logger

import (
	"log"
	"os"
)

// Log levels
const (
	InfoLevel = iota
	WarningLevel
	ErrorLevel
	FatalLevel
)

// Logger struct
type Logger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	fatalLogger   *log.Logger
}

var instance *Logger

// Init initializes the logger instance
func Init() {
	instance = &Logger{
		infoLogger:    log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warningLogger: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger:   log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		fatalLogger:   log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Log logs messages based on the provided log level
func Log(level int, msg string) {
	if instance == nil {
		Init()
	}
	switch level {
	case InfoLevel:
		instance.infoLogger.Println(msg)
	case WarningLevel:
		instance.warningLogger.Println(msg)
	case ErrorLevel:
		instance.errorLogger.Println(msg)
	case FatalLevel:
		instance.fatalLogger.Fatalln(msg)
	default:
		log.Println("Unknown log level:", msg)
	}
}
