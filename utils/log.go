package utils

import (
	"github.com/fatih/color"
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, color.New(color.FgBlue, color.Bold).Sprintf("[")+color.New(color.FgBlue).Sprintf("INFO")+color.New(color.FgBlue, color.Bold).Sprintf("] "), log.LstdFlags)
	warnLogger  = log.New(os.Stdout, color.New(color.FgYellow, color.Bold).Sprintf("[")+color.New(color.FgYellow).Sprintf("WARN")+color.New(color.FgYellow, color.Bold).Sprintf("] "), log.LstdFlags)
	errorLogger = log.New(os.Stderr, color.New(color.FgRed, color.Bold).Sprintf("[")+color.New(color.FgRed).Sprintf("ERROR")+color.New(color.FgRed, color.Bold).Sprintf("] "), log.LstdFlags)
)

func Info(message string, err error) {
	if err != nil {
		infoLogger.Printf("%s - Error: %v\n", message, err)
	} else {
		infoLogger.Println(message)
	}
}

func Warn(message string, err error) {
	if err != nil {
		warnLogger.Printf("%s - Error: %v\n", message, err)
	} else {
		warnLogger.Println(message)
	}
}

func Error(message string, err error) {
	if err != nil {
		errorLogger.Printf("%s - Error: %v\n", message, err)
	} else {
		errorLogger.Println(message)
	}
}
