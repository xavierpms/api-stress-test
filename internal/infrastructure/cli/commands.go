package cli

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:          "api-stress-test",
	Short:        "CLI tool for running HTTP stress tests and generating summary reports",
	RunE:         runStressTest(),
	SilenceUsage: true,
}

func init() {
	RootCmd.Flags().StringP("url", "u", "", "URL for service to be tested")
	RootCmd.Flags().IntP("requests", "r", 0, "Total number of requests")
	RootCmd.Flags().IntP("concurrency", "c", 1, "Number of concurrent requests")

	_ = RootCmd.MarkFlagRequired("url")
	_ = RootCmd.MarkFlagRequired("requests")
}
