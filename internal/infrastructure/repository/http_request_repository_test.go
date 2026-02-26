package repository

import (
	"errors"
	"net/http"
	"testing"
	"time"
)

type httpClientStub struct {
	status int
	err    error
}

func (c *httpClientStub) Do(_ *http.Request) (*http.Response, error) {
	if c.err != nil {
		return nil, c.err
	}
	return &http.Response{StatusCode: c.status, Body: http.NoBody}, nil
}

func TestRequestRepositoryDo(t *testing.T) {
	t.Run("creates one status per request", func(t *testing.T) {
		repository := NewRequestRepository(&httpClientStub{status: 200}, time.Second)

		statusCodes, _, err := repository.Do("http://example.com", 5, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(statusCodes) != 5 {
			t.Fatalf("expected 5 status codes, got %d", len(statusCodes))
		}
	})

	t.Run("maps transport errors to zero status code", func(t *testing.T) {
		repository := NewRequestRepository(&httpClientStub{err: errors.New("network")}, time.Second)

		statusCodes, _, err := repository.Do("http://example.com", 3, 1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		for _, code := range statusCodes {
			if code != 0 {
				t.Fatalf("expected status code 0 for transport error, got %d", code)
			}
		}
	})
}
