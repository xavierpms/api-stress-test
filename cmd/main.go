package main

import (
	"fmt"
	"os"

	"github.com/xavierpms/api-stress-test/internal/infrastructure/cli"
)

func main() {
	err := cli.RootCmd.Execute()
	if err != nil {
		fmt.Printf("Fail to execute root cmd: %v", err)
		os.Exit(1)
	}
}
