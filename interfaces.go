package main

import (
  "fmt"
)

type Logger interface {
  Log(message string)
}

type Server struct {
  logger Logger
}

func main() {
  server := ConsoleLogger{}
  server.Log("Hello!")
}

type ConsoleLogger struct {}
func (l ConsoleLogger) Log(message string) {
  fmt.Println(message)
}
