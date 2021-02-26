package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/l6p/utils/client/json"
	"time"
)

func SimpleCase(client *json.Client) {
	data := client.R().Get("https://jsonplaceholder.typicode.com/todos/1").D()

	// Show Responded Json
	println(data.GetJson(""))

	time.Sleep(1 * time.Second)
}

// You can debug the test cases by executing the main function locally.
func Export() map[string]interface{} {
	return map[string]interface{}{
		"SimpleCase": SimpleCase,
	}
}

func main() {
	client := json.NewClient(resty.New())
	SimpleCase(client)
}
