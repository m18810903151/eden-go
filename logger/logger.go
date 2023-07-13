package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	isDebug    = false
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// Setup initialize the log instance
func Setup() {
	logger = log.New(os.Stdout, DefaultPrefix, log.Ldate|log.Ltime|log.Lmicroseconds|log.Lmsgprefix)
}

func SetDebugMode(debug bool) {
	isDebug = debug
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	if !isDebug {
		return
	}
	setPrefix(DEBUG)
	logger.Println(v...)
}

// Info output logs at info level
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
}

// Error output logs at error level
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v...)
}

// setPrefix set the prefix of the log output
func setPrefix(level Level) {
	var logPrefix string
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("%s - %s:%d ", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("%s ", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}