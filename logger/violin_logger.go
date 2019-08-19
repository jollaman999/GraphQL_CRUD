package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

var Log_name = "violin"
var Logger *log.Logger
var once sync.Once
var FpLog *os.File

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func Prepare() bool {
	return_value := false

	once.Do(func() {
		// Check violin logger directory
		// Create directory if not exist
		if _, err := os.Stat("/var/log/" + Log_name); os.IsNotExist(err) {
			createDirIfNotExist("/var/log/" + Log_name)
		}

		now := time.Now()

		year := fmt.Sprintf("%d", now.Year())
		month := fmt.Sprintf("%02d", now.Month())
		day := fmt.Sprintf("%02d", now.Day())

		date := year + month + day

		var err error = nil
		FpLog, err = os.OpenFile("/var/log/" + Log_name + "/" +
			Log_name + "_" + date + ".log", os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}

		Logger = log.New(io.MultiWriter(FpLog, os.Stdout), "violin_logger: ", log.Ldate|log.Ltime)

		return_value = true
	})

	return return_value
}