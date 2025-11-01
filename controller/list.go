package controller

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	var diffCmd = &cobra.Command{
		Use:   "list",
		Short: "Show list of all environments",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("✅ Showing list of all environments\n")
		},
	}

	return diffCmd
}
