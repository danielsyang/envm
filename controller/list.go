package controller

import (
	"envm/service"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	var diffCmd = &cobra.Command{
		Use:   "list",
		Short: "Show list of all environments",
		Run: func(cmd *cobra.Command, args []string) {

			envs, err := service.ListEnvironments()

			if err != nil {
				fmt.Printf("❌ Error listing environments: %s\n", err)
				os.Exit(1)
			}

			fmt.Printf("List of environments:\n")
			for _, env := range envs {
				fmt.Printf("✅ %s\n", env)
			}

		},
	}

	return diffCmd
}
