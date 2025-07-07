package weverse

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"golang.org/x/net/proxy"
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

func MakeProxyClient(proxyURL string, timeout time.Duration) (*http.Client, error) {
	client := &http.Client{}
	if proxyURL != "" {
		dialer, err := proxy.SOCKS5("tcp", proxyURL, nil, proxy.Direct)
		if err != nil {
			return nil, err
		}
		transport := &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return dialer.Dial(network, addr)
			},
		}
		client = &http.Client{
			Transport: transport,
			Timeout:   timeout,
		}
	}
	return client, nil
}

func generateWeverseURL(targetPath string, queryParams map[string]string) (string, error) {
	if !strings.HasSuffix(targetPath, "?") {
		targetPath += "?"
	}
	wmsgpad := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	values := url.Values{}
	for k, v := range queryParams {
		values.Set(k, v)
	}
	encodedParams := values.Encode()
	apiPath := targetPath + encodedParams
	if len(apiPath) > 255 {
		apiPath = apiPath[:255]
	}
	mac := hmac.New(sha1.New, []byte(HMACKey))
	mac.Write([]byte(apiPath + wmsgpad))
	signature := mac.Sum(nil)
	wmd := base64.StdEncoding.EncodeToString(signature)
	finalQuery := encodedParams + "&wmsgpad=" + url.QueryEscape(wmsgpad) + "&wmd=" + url.QueryEscape(wmd)
	finalQuery = strings.TrimPrefix(finalQuery, "&")
	finalURL := weverseBaseURL + targetPath + finalQuery
	return finalURL, nil
}