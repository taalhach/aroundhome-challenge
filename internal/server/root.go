package server

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:  "challenge",
	Long: "bsdex code challenge",
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Printf("error while executing command: %v\n", err)
		os.Exit(1)
	}
}
