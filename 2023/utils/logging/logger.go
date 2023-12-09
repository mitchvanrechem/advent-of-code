package logging

import "fmt"

type Logger struct {
	enabled bool
}

type ILogger interface {
	Log(string)
}

func NewLogger(enabled bool) ILogger {
	return &Logger{enabled}
}

func (d *Logger) Log(message string) {
	if d.enabled {
		fmt.Println(message)
	}
}
