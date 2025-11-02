package controller

import (
	"envm/service"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func GenerateCmd() *cobra.Command {
	var generateCmd = &cobra.Command{
		Use:   "generate [environment]",
		Short: "build .env file from selected environment",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			env := args[0]

			content, err := service.GenerateEnvironment(env)

			if err != nil {
				fmt.Printf("❌ Error generating environment: %s\n", err)
				os.Exit(1)
			}

			// The third argument (0644) sets the file permissions (read/write for owner, read-only for others).
			err = os.WriteFile(".env", []byte(content), 0644)

			if err != nil {
				fmt.Printf("❌ Error writing .env file: %s\n", err)
				os.Exit(1)
			}

			fmt.Printf("✅ Generated .env file\n")
		},
	}

	return generateCmd
}
