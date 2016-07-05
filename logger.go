package goma

import (
	"log"
	"os"
	"runtime"
	"strings"
)

var info *log.Logger

func NewLogger() *log.Logger {
	info = log.New(os.Stdout, "GOMA: ", log.Lshortfile)
	info.Println("Logger inited.")
	return info
}

func Log(message interface{}) {
	// Deduce caller info
	_, file, no, _ := runtime.Caller(1)
	arr := strings.Split(file, "/")
	component := strings.Join(arr[len(arr)-2:len(arr)], "/")

	if info == nil {
		log.Println("Warning: Logger has not been inited. Please call NewLogger()")
	} else {
		info.Printf("%s:%d %v\n", component, no, message)
	}
}
