package application

import (
	"time"

	domain "github.com/xavierpms/api-stress-test/internal/domain"
)

type GenerateReportInputDTO struct {
	StatusCodes []int
	Duration    time.Duration
}

type GenerateReportOutputDTO struct {
	TimeSpent          string         `json:"time_spent"`
	RequestsMade       int            `json:"requests_made"`
	SuccessfulRequests int            `json:"successful_requests"`
	FailedRequests     map[string]int `json:"failed_requests"`
}

type GenerateReportUseCase struct {
	ReportRepository domain.ReportRepositoryInterface
}

func NewGenerateReportUseCase(reportRepository domain.ReportRepositoryInterface) *GenerateReportUseCase {
	return &GenerateReportUseCase{
		ReportRepository: reportRepository,
	}
}

func (r *GenerateReportUseCase) Execute(input GenerateReportInputDTO) (GenerateReportOutputDTO, error) {
	report, err := r.ReportRepository.Generate(input.StatusCodes, input.Duration)
	if err != nil {
		return GenerateReportOutputDTO{}, err
	}

	dto := GenerateReportOutputDTO{
		TimeSpent:          report.TimeSpent.String(),
		RequestsMade:       report.RequestsMade,
		SuccessfulRequests: report.SuccessfulRequests,
		FailedRequests:     report.FailedRequests,
	}

	return dto, nil
}
