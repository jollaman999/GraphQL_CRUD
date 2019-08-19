package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

var Log_name = "violin"
var Logger *log.Logger
var once sync.Once
var FpLog *os.File

func Prepare() bool {
	return_value := false

	once.Do(func() {
		// Check violin logger directory
		// Run mkdir if not exist
		if _, err := os.Stat("/var/log/" + Log_name); os.IsNotExist(err) {
			cmd := exec.Command("mkdir", "-p", "/var/log/" + Log_name + "/")
			err := cmd.Start()
			if err != nil {
				fmt.Println("Failed to create logger directory!")
				return_value = false
			}

			// Wait for create directory
			_ = cmd.Wait()
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