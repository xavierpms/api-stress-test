package application

import (
	"errors"
	"testing"
	"time"
)

type requestRepoStub struct {
	statusCodes []int
	duration    time.Duration
	err         error
}

func (s *requestRepoStub) Do(_ string, _, _ int) ([]int, time.Duration, error) {
	return s.statusCodes, s.duration, s.err
}

func TestDoRequestUseCaseExecute(t *testing.T) {
	t.Run("returns output from repository", func(t *testing.T) {
		repo := &requestRepoStub{statusCodes: []int{200, 500}, duration: 2 * time.Second}
		uc := NewDoRequestUseCase(repo)

		out, err := uc.Execute(DoRequestInputDTO{URL: "http://example.com", Requests: 2, Concurrency: 1})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(out.StatusCodes) != 2 {
			t.Fatalf("expected 2 status codes, got %d", len(out.StatusCodes))
		}
		if out.TimeSpent != 2*time.Second {
			t.Fatalf("expected duration 2s, got %s", out.TimeSpent)
		}
	})

	t.Run("propagates repository error", func(t *testing.T) {
		repo := &requestRepoStub{err: errors.New("boom")}
		uc := NewDoRequestUseCase(repo)

		_, err := uc.Execute(DoRequestInputDTO{URL: "http://example.com", Requests: 1, Concurrency: 1})
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}
