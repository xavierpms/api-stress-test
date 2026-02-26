package application

import (
	"errors"
	"testing"
	"time"
)

type doRequestExecutorStub struct {
	out DoRequestOutputDTO
	err error
}

func (s *doRequestExecutorStub) Execute(_ DoRequestInputDTO) (DoRequestOutputDTO, error) {
	return s.out, s.err
}

type reportExecutorStub struct {
	out GenerateReportOutputDTO
	err error
}

func (s *reportExecutorStub) Execute(_ GenerateReportInputDTO) (GenerateReportOutputDTO, error) {
	return s.out, s.err
}

func TestRunStressTestUseCaseExecute(t *testing.T) {
	t.Run("orchestrates request and report", func(t *testing.T) {
		runner := NewRunStressTestUseCase(
			&doRequestExecutorStub{out: DoRequestOutputDTO{StatusCodes: []int{200, 500}, TimeSpent: time.Second}},
			&reportExecutorStub{out: GenerateReportOutputDTO{RequestsMade: 2, SuccessfulRequests: 1}},
		)

		out, err := runner.Execute(RunStressTestInputDTO{URL: "http://example.com", Requests: 2, Concurrency: 1})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if out.RequestsMade != 2 {
			t.Fatalf("expected requests made 2, got %d", out.RequestsMade)
		}
	})

	t.Run("returns do request error", func(t *testing.T) {
		runner := NewRunStressTestUseCase(
			&doRequestExecutorStub{err: errors.New("request error")},
			&reportExecutorStub{},
		)

		_, err := runner.Execute(RunStressTestInputDTO{})
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("returns report error", func(t *testing.T) {
		runner := NewRunStressTestUseCase(
			&doRequestExecutorStub{out: DoRequestOutputDTO{StatusCodes: []int{200}, TimeSpent: time.Second}},
			&reportExecutorStub{err: errors.New("report error")},
		)

		_, err := runner.Execute(RunStressTestInputDTO{})
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}
