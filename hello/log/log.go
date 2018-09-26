package log

import "fmt"

type Logger interface {
  Log (message string)
}

// a simple logger
type SimpleLogger struct {
  prefix string // private members not visible outside this package
}

// constructor
func CreateSimpleLogger () *SimpleLogger {
  return &SimpleLogger { prefix: "simple", }
}

func (logger SimpleLogger) Log (message string) {
  fmt.Printf("[%s] %s\n", logger.prefix, message)
}

// a beatified(?) logger
type BeautyLogger struct {
}

// contructor
func CreateBeautyLogger () *BeautyLogger {
  return &BeautyLogger {}
}

func (logger BeautyLogger) Log (message string) {
  fmt.Printf("[beauty] %s\n", message)
}
