package cmd

import (
	"fmt"

	"veil/internals/workspace"

	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply <name>",
	Short: "Apply workspace changes back to the original project",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		ws, err := workspace.Load(name)
		if err != nil {
			fmt.Println("failed to load workspace:", err)
			return
		}

		if err := ws.Apply(); err != nil {
			fmt.Println("failed to apply changes:", err)
			return
		}

		fmt.Println()
		fmt.Println("VEIL   Changes applied")
		fmt.Println()
		fmt.Println("Workspace:", ws.Name)
		fmt.Println("Project:  ", ws.Project)
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}