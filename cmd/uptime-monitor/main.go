package main

import (
	"fmt"
	"time"

	"github.com/Symeks/uptime-monitor/internal/checker"
)

func main() {
	r := checker.Result{
		URL:          "https://example.com",
		CreatedAt:    time.Now(),
		StatusCode:   200,
		ResponseTime: 0,
		Success:      true,
		Err:          nil,
	}

	fmt.Printf("%+v\n", r)
}
