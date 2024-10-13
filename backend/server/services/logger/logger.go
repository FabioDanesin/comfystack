package logger

import "os"

type Logger struct {
	std *os.File
}

var Instance *Logger

func Create_default_logger() {
	Instance = &Logger{std: os.Stdout}
}

func Create_logger(std *os.File) {
	if std != nil {
		Instance = &Logger{std: std}
	}
}

func (l *Logger) LogWrite(msg string) {
	l.std.WriteString(msg)
}
