package logger

import "os"

type Logger struct {
	std *os.File
}

var Instance *Logger

func CreateDefaultLogger() {
	Instance = &Logger{std: os.Stdout}
}

func CreateLogger(std *os.File) {
	if std != nil {
		Instance = &Logger{std: std}
	}
}

func (l *Logger) LogWrite(msg string) {
	l.std.WriteString(msg)
}
