package main

import (
	"fmt"
	"github.com/l6p/utils/client/json"
	"math/rand"
	"time"
)

type Context struct {
	BaseUrl  string
	SubPaths []string `l6p:"sync"`
}

func SetUp(ctx *Context) {
	ctx.SubPaths = []string{
		"/users",
		"/todos",
		"/posts",
	}
}

func SimpleCase(ctx *Context, client *json.Client) {
	subPath := ctx.SubPaths[rand.Intn(len(ctx.SubPaths))]
	_ = client.R().Get(fmt.Sprintf("%s%s", ctx.BaseUrl, subPath))
	time.Sleep(5 * time.Second)
}

func Export() map[string]interface{} {
	return map[string]interface{}{
		"context": Context{
			BaseUrl: "https://jsonplaceholder.typicode.com",
		},
		"setUp":      SetUp,
		"SimpleCase": SimpleCase,
	}
}
