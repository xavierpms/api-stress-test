package repository

import (
	"time"

	domain "github.com/xavierpms/api-stress-test/internal/domain"
)

type ReportRepository struct{}

func NewReportRepository() *ReportRepository {
	return &ReportRepository{}
}

func (r *ReportRepository) Generate(statusCodes []int, duration time.Duration) (*domain.Report, error) {
	requestsMade := len(statusCodes)
	successfulRequests := 0
	failedRequests := map[string]int{
		"3xx":    0,
		"4xx":    0,
		"5xx":    0,
		"errors": 0,
	}

	for _, code := range statusCodes {
		if code == 200 {
			successfulRequests++
		} else {
			switch {
			case code >= 300 && code <= 399:
				failedRequests["3xx"]++
			case code >= 400 && code <= 499:
				failedRequests["4xx"]++
			case code >= 500 && code <= 599:
				failedRequests["5xx"]++
			default:
				failedRequests["errors"]++
			}
		}
	}

	report := domain.NewReport(duration, requestsMade, successfulRequests, failedRequests)

	return report, nil
}
