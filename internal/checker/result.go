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

	if err != nil {
		r.Err = err
		r.ResponseTime = time.Since(r.CreatedAt)
		r.Success = false
		return r
	}

	defer res.Body.Close()

	r.StatusCode = res.StatusCode
	r.ResponseTime = time.Since(r.CreatedAt)
	r.Success = checkStatus(res.StatusCode)

	return r
}

func checkStatus(statusCode int) bool {

	if statusCode < 200 || statusCode >= 300 {
		return false
	}

	return true
}
