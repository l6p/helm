package main

import (
	"fmt"
	"github.com/l6p/utils/client/json"
	"time"
)

type Context struct {
	BaseUrl string
}

func SimpleCase(ctx *Context, client *json.Client) {
	_ = client.R().Get(fmt.Sprintf("%s/todos/1", ctx.BaseUrl))
	time.Sleep(5 * time.Second)
}

func Export() map[string]interface{} {
	return map[string]interface{}{
		"context": Context{
			BaseUrl: "https://jsonplaceholder.typicode.com",
		},
		"SimpleCase": SimpleCase,
	}
}
