package domain

import (
	"time"
)

type RequestRepositoryInterface interface {
	Do(string, int, int) ([]int, time.Duration, error)
}

type ReportRepositoryInterface interface {
	Generate([]int, time.Duration) (*Report, error)
}
