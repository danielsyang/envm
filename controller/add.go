package controller

import (
	"fmt"

	"github.com/spf13/cobra"
)

func AddCmd() *cobra.Command {
	var addCmd = &cobra.Command{
		Use:   "add [environment]",
		Short: "Add a new active environment (local|dev|stage|prod)",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			env := args[0]
			fmt.Printf("✅ Environment set to: %s\n", env)
			// TODO: save to $HOME/.config/envm/state.json OR sqlite
		},
	}

	return addCmd
}
