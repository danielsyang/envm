package controller

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GenerateCmd() *cobra.Command {
	var generateCmd = &cobra.Command{
		Use:   "generate [environment]",
		Short: "build .env file from selected environment",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("✅ Generating .env file from layered config\n")
		},
	}

	return generateCmd
}
