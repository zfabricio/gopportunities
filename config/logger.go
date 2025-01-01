package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// Create Non-Formatted logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}
func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}
func (l *Logger) Err(v ...interface{}) {
	l.err.Println(v...)
}

// Create Format Enabled Logs
func (l *Logger) Debugf(f string, v ...interface{}) {
	l.debug.Printf(f, v...)
}
func (l *Logger) Infof(f string, v ...interface{}) {
	l.info.Printf(f, v...)
}
func (l *Logger) WarningF(f string, v ...interface{}) {
	l.warning.Printf(f, v...)
}
func (l *Logger) ErrF(f string, v ...interface{}) {
	l.err.Printf(f, v...)
}
