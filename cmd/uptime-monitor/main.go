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
		"https://httpbin.org/delay/3",
		"https://httpbin.org/delay/2",
	}

	t := time.Now()

	results := checker.CheckAll(client, urls)

	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}

	fmt.Printf("Total time taken: %v\n", time.Since(t))
}
