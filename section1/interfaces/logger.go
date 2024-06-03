package loggers

import (
	"errors"
	"fmt"
	"os"
)

type Logger interface {
	Log(message string)
}

type FileLogger struct {
	Destination string
}

type ConsoleLogger struct {
}

func (fileDetail *FileLogger) Log(message string) error {
	file, err := os.OpenFile(fileDetail.Destination, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return errors.New("error while opening file")
	}

	defer file.Close()

	_, err = file.WriteString(message)

	if err != nil {
		return errors.New("error while writing to file")
	}

	return nil
}

func (consoleDetail *ConsoleLogger) Log(message string) {
	fmt.Println(message)
}
