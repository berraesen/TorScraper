package scraper

import (
	"io"
	"net/http"
	"time"

	"golang.org/x/net/proxy"
)

func NewClient() (*http.Client, error) {

	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9150", nil, proxy.Direct)
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{
		Dial: dialer.Dial,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   20 * time.Second,
	}

	return client, nil
}

func TakeURL(client *http.Client, url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
