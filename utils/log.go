package utils

import (
	"github.com/fatih/color"
	"log"
	"os"
)

var (
	boldWhite = color.New(color.FgWhite, color.Bold).SprintFunc()

	infoLogger  = log.New(os.Stdout, color.New(color.FgBlue, color.Bold).Sprint("[INFO] "), log.LstdFlags)
	warnLogger  = log.New(os.Stdout, color.New(color.FgYellow, color.Bold).Sprint("[WARN] "), log.LstdFlags)
	errorLogger = log.New(os.Stderr, color.New(color.FgRed, color.Bold).Sprint("[ERROR] "), log.LstdFlags)
)

func Info(message string, err error) {
	if err != nil {
		infoLogger.Println(boldWhite(message + " - Error: " + err.Error()))
	} else {
		infoLogger.Println(boldWhite(message))
	}
}

func Warn(message string, err error) {
	if err != nil {
		warnLogger.Println(boldWhite(message + " - Error: " + err.Error()))
	} else {
		warnLogger.Println(boldWhite(message))
	}
}

func Error(message string, err error) {
	if err != nil {
		errorLogger.Println(boldWhite(message + " - Error: " + err.Error()))
	} else {
		errorLogger.Println(boldWhite(message))
	}
}
