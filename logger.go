package logger

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	AllLevel
)

type Logger struct {
	*log.Logger
	level         LogLevel
	logFormatFlag int
}

func New() *Logger {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	logFileName := os.Getenv("LOG_FILE")
	fmt.Println("logfile name printed here: ", logFileName)
	if logFileName == "" {
		logFileName = "logfile.log"
	}

	logLevelStr := os.Getenv("LOG_LEVEL")
	logLevel := LogLevelFromString(logLevelStr)

	logFormatStr := os.Getenv("LOG_FORMAT")
	logFormatFlag := LogFormatToFlag(logFormatStr)

	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	return &Logger{
		Logger:        log.New(file, "Logger: ", log.Ldate|log.Ltime|log.Lshortfile),
		level:         logLevel,
		logFormatFlag: logFormatFlag,
	}
}

func (cl *Logger) Debug(msg string) {
	if cl.level <= DebugLevel || cl.level == AllLevel {
		cl.Println("[DEBUG]", msg)
	}
}

func (cl *Logger) Info(msg string) {
	if cl.level <= InfoLevel || cl.level == AllLevel {
		cl.Println("[INFO]", msg)
	}
}

func (cl *Logger) Warning(msg string) {
	if cl.level <= WarningLevel || cl.level == AllLevel {
		cl.Println("[WARNING]", msg)
	}
}

func (cl *Logger) Error(msg string) {
	if cl.level <= ErrorLevel || cl.level == AllLevel {
		cl.Println("[ERROR]", msg)
	}
}

func LogLevelFromString(levelStr string) LogLevel {
	switch levelStr {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warning":
		return WarningLevel
	case "error":
		return ErrorLevel
	default:
		return AllLevel // Default log level
	}
}

func LogFormatToFlag(formatStr string) int {
	// Split the format string by "|"
	formatParts := strings.Split(formatStr, "|")

	// Initialize logFlag with 0
	logFlag := 0

	// Parse each format part and add corresponding flag to logFlag
	for _, part := range formatParts {
		switch strings.TrimSpace(part) {
		case "date":
			logFlag |= log.Ldate
		case "time":
			logFlag |= log.Ltime
		case "microseconds":
			logFlag |= log.Lmicroseconds
		case "longfile":
			logFlag |= log.Llongfile
		case "shortfile":
			logFlag |= log.Lshortfile
		case "utc":
			logFlag |= log.LUTC
		case "stdflags":
			logFlag |= log.LstdFlags
		}
	}

	return logFlag
}
