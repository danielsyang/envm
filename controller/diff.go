package controller

import (
	"fmt"

	"github.com/spf13/cobra"
)

func DiffCmd() *cobra.Command {
	var diffCmd = &cobra.Command{
		Use:   "diff [environment1] [environment2]",
		Short: "Show key differences between two envs",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			env1 := args[0]
			env2 := args[1]
			fmt.Printf("✅ Showing key differences between %s and %s\n", env1, env2)
		},
	}

	return diffCmd
}
