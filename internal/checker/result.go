package checker

import (
	"net/http"
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

func CheckURL(client *http.Client, url string) Result {
	r := Result{
		URL:       url,
		CreatedAt: time.Now(),
	}

	res, err := client.Get(url)
	rTime := time.Since(r.CreatedAt)

	if err != nil {
		r.Err = err
		r.ResponseTime = rTime
		r.Success = false
		return r
	}

	defer res.Body.Close()

	r.StatusCode = res.StatusCode
	r.ResponseTime = rTime
	r.Success = true

	return r
}
