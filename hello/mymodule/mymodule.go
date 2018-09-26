package mymodule

import "fmt"
import "hello/log"

func Show (id int) {
  show(id)
}

// this function is not visible outside of the package
func show (id int) {
  fmt.Printf("Mymodule show %d\n", id)
}

// this function accepts a log interface to show a message
func ShowLog(logger log.Logger, message string) {
  logger.Log(message)
}
