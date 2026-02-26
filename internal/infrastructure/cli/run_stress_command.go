package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	app "github.com/xavierpms/api-stress-test/internal/application"
	repository "github.com/xavierpms/api-stress-test/internal/infrastructure/repository"
)

type RunEFunc func(cmd *cobra.Command, args []string) error

func runStressTest() RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		input, err := parseInput(cmd)
		if err != nil {
			return err
		}

		requestRepo := repository.NewDefaultRequestRepository()
		reportRepo := repository.NewReportRepository()

		runner := app.NewRunStressTestUseCase(
			app.NewDoRequestUseCase(requestRepo),
			app.NewGenerateReportUseCase(reportRepo),
		)

		report, err := runner.Execute(input)
		if err != nil {
			return err
		}

		report_output, err := json.MarshalIndent(report, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(report_output))

		return nil
	}
}

func parseInput(cmd *cobra.Command) (app.RunStressTestInputDTO, error) {
	url, err := cmd.Flags().GetString("url")
	if err != nil {
		return app.RunStressTestInputDTO{}, err
	}
	url = strings.TrimSpace(url)
	if url == "" {
		return app.RunStressTestInputDTO{}, errors.New("invalid URL")
	}

	requests, err := cmd.Flags().GetInt("requests")
	if err != nil {
		return app.RunStressTestInputDTO{}, err
	}
	if requests <= 0 {
		return app.RunStressTestInputDTO{}, errors.New("invalid requests")
	}

	concurrency, err := cmd.Flags().GetInt("concurrency")
	if err != nil {
		return app.RunStressTestInputDTO{}, err
	}
	if concurrency <= 0 {
		concurrency = 1
	}

	return app.RunStressTestInputDTO{
		URL:         url,
		Requests:    requests,
		Concurrency: concurrency,
	}, nil
}
