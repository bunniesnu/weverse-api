package weverse

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

// clickLink performs an HTTP GET to the given link,
// mimicking a browser click by sending common headers,
// handling cookies, and following redirects.
// It returns the final response body as a string.
func clickLink(rawURL string) error {
    ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var html string
	err := chromedp.Run(ctx,
		chromedp.Navigate(rawURL),
		chromedp.Sleep(5*time.Second), // wait for JS
		chromedp.OuterHTML("html", &html),
	)
	return err
}