package logger

import (
	"fmt"
	"log"
	"os"
)

type (
	Level int
)

const (
	LevelFatal = iota
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
	fatalColor   = "\033[1;37;41m"
	errColor     = "\033[1;37;45m"
	warningColor = "\033[1;37;43m"
	infoColor    = "\033[1;37;42m"
	mColor       = "\033[1;37;44m"
	debugColor   = "\033[1;36m"
	colorEnd     = "\033[0m"
	entName      = "DQK"
)

var logger = New()

func p(level Level, v ...interface{}) {
	var formatStr string
	length := len(v)
	for i := 0; i < length; i++ {
		formatStr += "%+v "
	}
	logger.Output(level, fmt.Sprintf(formatStr, v...))
}

func Fatal(v ...interface{}) {
	p(LevelFatal, v...)
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	logger.Output(LevelFatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Error(v ...interface{}) {
	p(LevelError, v...)
}

func Errorf(format string, v ...interface{}) {
	logger.Output(LevelError, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	p(LevelWarning, v...)
}

func Warnf(format string, v ...interface{}) {
	logger.Output(LevelWarning, fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	p(LevelInfo, v...)
}

func Infof(format string, v ...interface{}) {
	logger.Output(LevelInfo, fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
	p(LevelDebug, v...)
}

func Debugf(format string, v ...interface{}) {
	logger.Output(LevelDebug, fmt.Sprintf(format, v...))
}

func Check(s string, err error) {
	if err == nil {
		logger.Output(LevelInfo, s+" ok")
	} else {
		logger.Output(LevelFatal, fmt.Sprintf("%s %+v", s, err))
		os.Exit(1)
	}
}

func SetLogLevel(level Level) {
	logger.SetLogLevel(level)
}

type logManager struct {
	_log *log.Logger
	//小于等于该级别的level才会被记录
	logLevel Level
}

//NewLogger 实例化，供自定义
func NewLogger() *logManager {
	return &logManager{_log: log.New(os.Stderr, "", log.Lshortfile|log.LstdFlags), logLevel: LevelDebug}
}

//New 实例化，供外部直接调用 log.XXXX
func New() *logManager {
	//return &logManager{_log: log.New(os.Stderr, "", log.Lshortfile|log.LstdFlags), logLevel: LevelDebug}
	return &logManager{_log: log.New(os.Stdout, fmt.Sprintf("[%s%s%s] ", mColor, entName, colorEnd), log.Lshortfile|log.LstdFlags), logLevel: LevelDebug}
}

func (l *logManager) Output(level Level, s string) error {
	if l.logLevel < level {
		return nil
	}
	switch level {
	case LevelFatal:
		s = fmt.Sprintf("[%sFATAL%s] %s", fatalColor, colorEnd, s)
	case LevelError:
		s = fmt.Sprintf("[%sERROR%s] %s", errColor, colorEnd, s)
	case LevelWarning:
		s = fmt.Sprintf("[%sWARNING%s] %s", warningColor, colorEnd, s)
	case LevelInfo:
		s = fmt.Sprintf("[%sINFO%s] %s", infoColor, colorEnd, s)
	case LevelDebug:
		s = fmt.Sprintf("[%sDEBUG%s] %s", debugColor, colorEnd, s)
	default:
		s = fmt.Sprintf("[%sUNKNOWN%s] %s", infoColor, colorEnd, s)
	}
	return l._log.Output(3, s)
}

func (l *logManager) Fatal(s string) {
	l.Output(LevelFatal, s)
	os.Exit(1)
}

func (l *logManager) Fatalf(format string, v ...interface{}) {
	l.Output(LevelFatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *logManager) Error(s string) {
	l.Output(LevelError, s)
}

func (l *logManager) Errorf(format string, v ...interface{}) {
	l.Output(LevelError, fmt.Sprintf(format, v...))
}

func (l *logManager) Warn(s string) {
	l.Output(LevelWarning, s)
}

func (l *logManager) Warnf(format string, v ...interface{}) {
	l.Output(LevelWarning, fmt.Sprintf(format, v...))
}

func (l *logManager) Info(s string) {
	l.Output(LevelInfo, s)
}

func (l *logManager) Infof(format string, v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *logManager) Debug(s string) {
	l.Output(LevelDebug, s)
}

func (l *logManager) Debugf(format string, v ...interface{}) {
	l.Output(LevelDebug, fmt.Sprintf(format, v...))
}

func (l *logManager) Check(s string, err error) {
	if err != nil {
		l.Output(LevelInfo, s+"ok")
	} else {
		l.Output(LevelWarning, s+fmt.Sprintf("%+v", err))
	}
}

func (l *logManager) SetLogLevel(level Level) {
	l.logLevel = level
}
