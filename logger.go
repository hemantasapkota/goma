package goma

import (
	"log"
	"os"
	"runtime"
	"strings"
)

type LoggerConfig struct {
	Debug bool
}

var info *log.Logger
var loggerConfig LoggerConfig

func NewLogger(conf LoggerConfig) *log.Logger {
	loggerConfig = conf
	info = log.New(os.Stdout, "GOMA: ", log.Lshortfile)
	info.Println("Logger inited.")
	return info
}

func Log(message interface{}) {
	if info == nil {
		log.Println("Warning: Logger has not been inited. Please call NewLogger()")
		return
	}

	if !loggerConfig.Debug {
		return
	}

	// Deduce caller info
	_, file, no, _ := runtime.Caller(1)
	arr := strings.Split(file, "/")
	component := strings.Join(arr[len(arr)-2:len(arr)], "/")

	info.Printf("%s:%d %v\n", component, no, message)
}
