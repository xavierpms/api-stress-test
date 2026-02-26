package application

import (
	"errors"
	"testing"
	"time"

	domain "github.com/xavierpms/api-stress-test/internal/domain"
)

type reportRepoStub struct {
	report *domain.Report
	err    error
}

func (s *reportRepoStub) Generate(_ []int, _ time.Duration) (*domain.Report, error) {
	return s.report, s.err
}

func TestGenerateReportUseCaseExecute(t *testing.T) {
	t.Run("maps entity report to output dto", func(t *testing.T) {
		repo := &reportRepoStub{report: domain.NewReport(3*time.Second, 10, 7, map[string]int{"4xx": 3})}
		uc := NewGenerateReportUseCase(repo)

		out, err := uc.Execute(GenerateReportInputDTO{StatusCodes: []int{200}, Duration: time.Second})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if out.RequestsMade != 10 {
			t.Fatalf("expected requests made 10, got %d", out.RequestsMade)
		}
		if out.SuccessfulRequests != 7 {
			t.Fatalf("expected successful requests 7, got %d", out.SuccessfulRequests)
		}
	})

	t.Run("propagates repository error", func(t *testing.T) {
		repo := &reportRepoStub{err: errors.New("report error")}
		uc := NewGenerateReportUseCase(repo)

		_, err := uc.Execute(GenerateReportInputDTO{})
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}
