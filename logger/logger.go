package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var (
	enabledFileLog = true
	logFilename    = "log.txt"
	mu             sync.Mutex
)

// Log writes text to standard output and file
func Log(text string) {
	log.Print(text)

	if !enabledFileLog {
		return
	}
	mu.Lock()
	defer mu.Unlock()
	if err := writeToFile(logFilename, text); err != nil {
		log.Println(err)
	}
}

// Logln writes text with new line to standard output and file
func Logln(v ...interface{}) {
	Log(fmt.Sprintln(v...))
}

// Logf writes formated text to standard output and file
func Logf(format string, v ...interface{}) {
	Log(fmt.Sprintf(format, v...))
}

// SetFilename updates filename in which logs will be saved
func SetFilename(fileName string) {
	logFilename = fileName
}

// Disable logging to file
func Disable() {
	enabledFileLog = false
}

// Enable logging to file
func Enable() {
	enabledFileLog = true
}

// writeToFile writes text to fileName, creates new one if it doesn't exist
func writeToFile(fileName, text string) error {
	file, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	text = fmt.Sprintf("%s %s", time.Now().String(), text)
	if _, err = file.WriteString(text); err != nil {
		return err
	}
	return nil
}
