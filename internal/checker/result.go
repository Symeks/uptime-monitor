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

func CheckAll(client *http.Client, urls []string) []Result {
	ch := make(chan Result, len(urls))

	for _, url := range urls {
		go func() {
			ch <- CheckURL(client, url)
		}()
	}

	results := make([]Result, len(urls))
	for i := 0; i < len(urls); i++ {
		results[i] = <-ch
	}
	return results
}
