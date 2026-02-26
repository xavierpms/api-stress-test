package repository

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RequestRepository struct {
	client  HTTPClient
	timeout time.Duration
}

func NewRequestRepository(client HTTPClient, timeout time.Duration) *RequestRepository {
	if timeout <= 0 {
		timeout = 10 * time.Second
	}

	if client == nil {
		client = http.DefaultClient
	}

	return &RequestRepository{client: client, timeout: timeout}
}

func NewDefaultRequestRepository() *RequestRepository {
	return NewRequestRepository(http.DefaultClient, 10*time.Second)
}

func (r *RequestRepository) doRequest(url string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, err
	}

	res, err := r.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	return res.StatusCode, nil
}

func (r *RequestRepository) Do(url string, requests, concurrency int) ([]int, time.Duration, error) {
	if concurrency <= 0 {
		concurrency = 1
	}

	statusCodeList := make([]int, 0, requests)
	var mu sync.Mutex
	var wg sync.WaitGroup
	jobs := make(chan struct{}, requests)

	start := time.Now()

	for worker := 0; worker < concurrency; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for range jobs {
				statusCode, err := r.doRequest(url)
				if err != nil {
					statusCode = 0
				}

				mu.Lock()
				statusCodeList = append(statusCodeList, statusCode)
				mu.Unlock()
			}
		}()
	}

	for i := 0; i < requests; i++ {
		jobs <- struct{}{}
	}
	close(jobs)

	wg.Wait()
	duration := time.Since(start)

	return statusCodeList, duration, nil
}
