package controller

import (
	"envm/service"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func AddCmd() *cobra.Command {
	var addCmd = &cobra.Command{
		Use:   "add [environment]",
		Short: "Add a new active environment (local|dev|stage|prod)",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			env := args[0]
			fileName := args[1]

			file, err := os.ReadFile(fileName)

			if err != nil {
				fmt.Printf("❌ Error reading file: %s\n", err)
				os.Exit(1)
			}

			fmt.Printf("✅ File content: %s\n", string(file))

			err = service.AddEnvironment(env, file)

			if err != nil {
				fmt.Printf("❌ Error adding environment: %s\n", err)
				os.Exit(1)
			}

			fmt.Printf("✅ Environment added successfully\n")
		},
	}

	return addCmd
}
