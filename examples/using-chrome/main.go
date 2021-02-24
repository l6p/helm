package main

import (
	"github.com/l6p/utils/client/web"
	"log"
	"time"
)

func SimpleCase(client *web.Client, logger *log.Logger) {
	var example string
	client.Go(`https://golang.org/pkg/time/`).
		WaitVisible(`body > footer`).
		Click(`#pkg-examples > div`).
		Value(`#example_After .play .input textarea`, &example).
		Do()
	logger.Printf("Go's time.After example:\n%s", example)
	time.Sleep(5 * time.Second)
}

func Export() map[string]interface{} {
	return map[string]interface{}{
		"SimpleCase": SimpleCase,
	}
}
