package application

type RunStressTestInputDTO struct {
	URL         string
	Requests    int
	Concurrency int
}

type DoRequestExecutor interface {
	Execute(input DoRequestInputDTO) (DoRequestOutputDTO, error)
}

type ReportExecutor interface {
	Execute(input GenerateReportInputDTO) (GenerateReportOutputDTO, error)
}

type RunStressTestUseCase struct {
	doRequestUseCase      DoRequestExecutor
	generateReportUseCase ReportExecutor
}

func NewRunStressTestUseCase(
	doRequestUseCase DoRequestExecutor,
	generateReportUseCase ReportExecutor,
) *RunStressTestUseCase {
	return &RunStressTestUseCase{
		doRequestUseCase:      doRequestUseCase,
		generateReportUseCase: generateReportUseCase,
	}
}

func (u *RunStressTestUseCase) Execute(input RunStressTestInputDTO) (GenerateReportOutputDTO, error) {
	requestOutput, err := u.doRequestUseCase.Execute(DoRequestInputDTO{
		URL:         input.URL,
		Requests:    input.Requests,
		Concurrency: input.Concurrency,
	})
	if err != nil {
		return GenerateReportOutputDTO{}, err
	}

	report, err := u.generateReportUseCase.Execute(GenerateReportInputDTO{
		StatusCodes: requestOutput.StatusCodes,
		Duration:    requestOutput.TimeSpent,
	})
	if err != nil {
		return GenerateReportOutputDTO{}, err
	}

	return report, nil
}
