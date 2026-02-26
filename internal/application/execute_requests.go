package application

import (
	"time"

	domain "github.com/xavierpms/api-stress-test/internal/domain"
)

type DoRequestInputDTO struct {
	URL         string
	Requests    int
	Concurrency int
}

type DoRequestOutputDTO struct {
	StatusCodes []int
	TimeSpent   time.Duration
}

type DoRequestUseCase struct {
	RequestRepository domain.RequestRepositoryInterface
}

func NewDoRequestUseCase(requestRepository domain.RequestRepositoryInterface) *DoRequestUseCase {
	return &DoRequestUseCase{
		RequestRepository: requestRepository,
	}
}

func (r *DoRequestUseCase) Execute(input DoRequestInputDTO) (DoRequestOutputDTO, error) {
	statusCodes, timeSpent, err := r.RequestRepository.Do(input.URL, input.Requests, input.Concurrency)
	if err != nil {
		return DoRequestOutputDTO{}, err
	}

	dto := DoRequestOutputDTO{
		StatusCodes: statusCodes,
		TimeSpent:   timeSpent,
	}

	return dto, nil
}
