package loggo

import (
	"log"
	"os"
)

type Logger interface {
	Error(output ...any)
	Errorf(format string, output ...any)
	Info(output ...any)
	Infof(format string, output ...any)
	Warn(output ...any)
	Warnf(format string, output ...any)
	Fatal(output ...any)
	Fatalf(format string, output ...any)
	Panic(output ...any)
	Close()
	GetLogFile() *os.File
}

type logger struct {
	*log.Logger
	file *os.File
}

func New() Logger {
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	return &logger{
		Logger: log.New(file, "", log.LstdFlags),
		file:   file,
	}
}

// Error implements Logger
func (l *logger) Error(output ...any) {
	st := []any{"[ERROR]"}
	st = append(st, output...)
	l.Println(st...)
	log.Println(st...)
}

// Errorf implements Logger
func (l *logger) Errorf(format string, output ...any) {
	l.Printf("[ERROR]"+format, output...)
	log.Printf("[ERROR]"+format, output...)
}

// Info implements Logger
func (l *logger) Info(output ...any) {
	st := []any{"[INFO]"}
	st = append(st, output...)
	l.Println(st...)
	log.Println(st...)
}

// Infof implements Logger
func (l *logger) Infof(format string, output ...any) {
	l.Printf("[INFO]"+format, output...)
	log.Printf("[INFO]"+format, output...)
}

// Warn implements Logger
func (l *logger) Warn(output ...any) {
	st := []any{"[WARNING]"}
	st = append(st, output...)
	l.Println(st...)
	log.Println(st...)
}

// Warnf implements Logger
func (l *logger) Warnf(format string, output ...any) {
	l.Printf("[WARNING]"+format, output...)
	log.Printf("[WARNING]"+format, output...)
}

// Fatal implements Logger
func (l *logger) Fatal(output ...any) {
	st := []any{"[FATAL]"}
	st = append(st, output...)
	log.Println(st...)
	l.Fatalln(st...)
}

// Fatalf implements Logger
func (l *logger) Fatalf(format string, output ...any) {
	l.Printf("[FATAL]"+format, output...)
	log.Fatalf("[WARNING]"+format, output...)
}

// Panic implements Logger
func (l *logger) Panic(output ...any) {
	st := []any{"[PANIC]"}
	st = append(st, output...)
	log.Println(st...)
	log.Panicln(st...)
}

func (l *logger) Close() {
	l.file.Close()
}

func (l *logger) GetLogFile() *os.File {
	return l.file
}
