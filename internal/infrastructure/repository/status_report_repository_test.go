package repository

import (
	"testing"
	"time"
)

func TestReportRepositoryGenerate(t *testing.T) {
	repository := NewReportRepository()

	report, err := repository.Generate([]int{200, 302, 404, 500, 0}, 2*time.Second)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if report.RequestsMade != 5 {
		t.Fatalf("expected requests made 5, got %d", report.RequestsMade)
	}
	if report.SuccessfulRequests != 1 {
		t.Fatalf("expected successful requests 1, got %d", report.SuccessfulRequests)
	}
	if report.FailedRequests["3xx"] != 1 || report.FailedRequests["4xx"] != 1 || report.FailedRequests["5xx"] != 1 || report.FailedRequests["errors"] != 1 {
		t.Fatalf("unexpected failed requests map: %+v", report.FailedRequests)
	}
}
