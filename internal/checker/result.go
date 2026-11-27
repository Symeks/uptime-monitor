package checker

import (
	"time"
)

type Result struct {
	URL          string
	CreatedAt    time.Time
	StatusCode   int
	ResponseTime time.Duration
	Success      bool
	Err          error
}
