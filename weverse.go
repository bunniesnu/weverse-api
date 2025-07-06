package weverse

import (
	"context"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/proxy"
)

type Weverse struct {
	Client *http.Client
	Email string
	Password string
	Nickname string
	AccessToken string
}

func New(email, password, proxyURL string, timeout time.Duration) (*Weverse, error) {
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
	return &Weverse{
		Client: client,
		Email: email,
		Password: password,
		Nickname: "",
		AccessToken: "",
	}, nil
}