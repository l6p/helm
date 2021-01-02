package main

import (
	"context"
	"github.com/chromedp/chromedp"
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

func main() {
	opts := []chromedp.ExecAllocatorOption{
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"),
		chromedp.WindowSize(1920, 1080),
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
	}

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	client := web.NewClient(&ctx)

	logger := log.New(log.Writer(), "", 0)
	SimpleCase(client, logger)
}
