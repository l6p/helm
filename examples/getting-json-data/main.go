package main

import (
	"github.com/l6p/utils/client/json"
	"time"
)

func SimpleCase(client *json.Client) {
	_ = client.R().Get("https://jsonplaceholder.typicode.com/todos/1")
	time.Sleep(5 * time.Second)
}

func Export() map[string]interface{} {
	return map[string]interface{}{
		"SimpleCase": SimpleCase,
	}
}
