package mylib

import (
	"log"
	"os"
)

func Log() {
	file, err := os.Create("logfile.log")
	if err != nil {
		log.Fatal("Cannot create log file:", err)
	}
	defer file.Close()
	log.SetOutput(file)

	// Log messages
	log.Println("This is a log message.")
	log.Printf("This is a formatted log message with value %d.", 42)
}
