package cmd

import (
	"fmt"

	"veil/internals/workspace"

	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy <name>",
	Short: "Destroy a Veil workspace",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		ws, err := workspace.Load(name)
		if err != nil {
			fmt.Println("failed to load workspace:", err)
			return
		}

		if err := ws.Destroy(); err != nil {
			fmt.Println("failed to destroy workspace:", err)
			return
		}

		fmt.Println()
		fmt.Println("VEIL   Workspace destroyed")
		fmt.Println()
		fmt.Println("Name:", ws.Name)
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)
}