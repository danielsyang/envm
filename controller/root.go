package controller

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "envm",
		Short: "A CLI tool for environment management",
		Long:  `envm — simple environment composer CLI`,
	}

	rootCmd.AddCommand(AddCmd())
	rootCmd.AddCommand(GenerateCmd())
	rootCmd.AddCommand(DiffCmd())
	rootCmd.AddCommand(ListCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
