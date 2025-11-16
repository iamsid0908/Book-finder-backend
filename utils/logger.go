package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

// type for log severity level
type LogLevel string
type color string

const (
	INFO     LogLevel = "INFO"     // info level log
	WARNING  LogLevel = "WARNING"  // warning level log
	ERROR    LogLevel = "ERROR"    // error level log
	REQUEST  LogLevel = "REQUEST"  // request level log
	RESPONSE LogLevel = "RESPONSE" // response level log
	GREEN    color    = "\033[32m"
	YELLOW   color    = "\033[33m"
	RED      color    = "\033[31m"
	CYAN     color    = "\033[36m"
	RESET    color    = "\033[0m"
)

type logData struct {
	Level          LogLevel    `json:"level,omitempty"`
	Timestamp      time.Time   `json:"timestamp,omitempty"`
	Caller         string      `json:"caller,omitempty"`
	Message        string      `json:"message,omitempty"`
	AdditionalInfo interface{} `json:"additionalInfo,omitempty"`
}

type logColor struct {
	Start color
	Reset color
}

var levelColor = map[LogLevel]logColor{
	INFO: {
		Start: GREEN,
		Reset: RESET,
	},
	WARNING: {
		Start: YELLOW,
		Reset: RESET,
	},
	ERROR: {
		Start: RED,
		Reset: RESET,
	},
	REQUEST: {
		Start: CYAN,
		Reset: RESET,
	},
	RESPONSE: {
		Start: CYAN,
		Reset: RESET,
	},
}

func setColor(level LogLevel) string {
	var (
		color    = levelColor[level]
		levelStr = fmt.Sprintf("[%v]", level)
	)

	return fmt.Sprint(color.Start, levelStr, color.Reset)
}

func logger(level LogLevel, message string, additionalInfo map[string]interface{}) {
	header := setColor(level)
	_, file, line, _ := runtime.Caller(2)
	funcName := fmt.Sprintf("%s:%d", file, line)

	currentTime := time.Now()
	newLog := logData{
		Level:          level,
		Timestamp:      currentTime,
		Caller:         funcName,
		Message:        message,
		AdditionalInfo: additionalInfo,
	}

	logByteInJson, err := json.MarshalIndent(newLog, "", "	")
	if err != nil {
		log.Println(header)
		fmt.Println(err.Error())
	}

	logByteInString, err := json.Marshal(newLog)
	if err != nil {
		log.Println(header)
		fmt.Println(err.Error())
	}

	logStringInJson := string(logByteInJson)
	logStringInString := string(logByteInString)
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(header)
		fmt.Println(err.Error())
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	fmt.Printf("%s\n%s", header, logStringInJson)
	log.SetFlags(0)
	log.Print(logStringInString)
}

// Generate error logger on terminal
func LogError(err error, additionalInfo map[string]interface{}) error {
	level := ERROR
	logger(level, err.Error(), additionalInfo)
	return err
}

// Generate warning logger on terminal
func LogWarning(message string, additionalInfo map[string]interface{}) {
	level := WARNING
	logger(level, message, additionalInfo)
}

// Generate info logger on terminal
func LogInfo(message string, additionalInfo map[string]interface{}) {
	level := INFO
	logger(level, message, additionalInfo)
}
