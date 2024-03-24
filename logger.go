package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
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

func (l *Logger) Debug(msg string, args ...interface{}) {
	if l.level <= DebugLevel || l.level == AllLevel {
		l.printLog("[DEBUG]", msg, args...)
	}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	if l.level <= InfoLevel || l.level == AllLevel {
		l.printLog("[INFO]", msg, args...)
	}
}

func (l *Logger) Warning(msg string, args ...interface{}) {
	if l.level <= WarningLevel || l.level == AllLevel {
		l.printLog("[WARNING]", msg, args...)
	}
}

func (l *Logger) Error(msg string, args ...interface{}) {
	if l.level <= ErrorLevel || l.level == AllLevel {
		l.printLog("[ERROR]", msg, args...)
	}
}

func (l *Logger) printLog(level, formattedMsg string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	file = trimGOPATH(file)

	msg := fmt.Sprintf(formattedMsg, args...)

	l.Printf("%s %s:%d %s", level, file, line, msg)
}

func trimGOPATH(path string) string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return path
	}
	return strings.TrimPrefix(path, gopath+"/src/")
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
