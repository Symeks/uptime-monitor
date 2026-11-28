package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Symeks/uptime-monitor/internal/checker"
)

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	urls := []string{
		"https://example.com",
		"https://httpbin.org/status/404",
		"https://diese-domain-gibt-es-nicht-xyz.de",
		"https://httpbin.org/status/500",
		"http://github.com",
	}

	for _, url := range urls {
		result := checker.CheckURL(client, url)
		fmt.Printf("%+v\n", result)
	}
}
