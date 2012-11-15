package logger

import "fmt"

type Logger struct {
	category string
}

func New(category string) *Logger {
	var l = Logger{category: category}
	return &l
}

func (l *Logger) Log(msg ...interface{}) {
    o := fmt.Sprintf(":%s:", l.category)
    msg = append([]interface{}{o}, msg...)
	fmt.Println(msg...)
}

func (l *Logger) SetCategory(category string) {
	l.category = category
}

func (l *Logger) Category() string {
	return l.category
}
