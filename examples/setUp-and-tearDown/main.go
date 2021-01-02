package main

import (
	"github.com/l6p/utils/client/json"
	"log"
	"time"
)

func SetUp(logger *log.Logger) {
	logger.Print("SetUp has been executed.")
}

func TearDown(logger *log.Logger) {
	logger.Print("TearDown has been executed.")
}

func SimpleCase(client *json.Client) {
	_ = client.R().Get("https://jsonplaceholder.typicode.com/todos/1")
	time.Sleep(5 * time.Second)
}

func Export() map[string]interface{} {
	return map[string]interface{}{
		"setUp":      SetUp,
		"tearDown":   TearDown,
		"SimpleCase": SimpleCase,
	}
}
