package main

import (
	"github.com/l6p/utils/client/json"
	"time"
)

type Context struct {
	Url     string `l6p:"url"`
	Timeout int    `l6p:"timeout"`
}

func SimpleCase(ctx *Context, client *json.Client) {
	_ = client.R().Get(ctx.Url)
	time.Sleep(time.Duration(ctx.Timeout) * time.Second)
}

func Export() map[string]interface{} {
	return map[string]interface{}{
		"context":    Context{},
		"SimpleCase": SimpleCase,
	}
}
