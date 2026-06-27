package cmd

import (
	"fmt"

	"veil/internals/workspace"

	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect <name>",
	Short: "Inspect a Veil workspace",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		ws, err := workspace.Load(args[0])
		if err != nil {
			fmt.Println("failed to load workspace:", err)
			return
		}

		status := "unmounted"
		if ws.IsMounted() {
			status = "mounted"
		}

		fmt.Println()
		fmt.Println("VEIL   Workspace Inspect")
		fmt.Println()
		fmt.Println("Name:   ", ws.Name)
		fmt.Println("Project:", ws.Project)
		fmt.Println("Upper:  ", ws.Upper)
		fmt.Println("Work:   ", ws.Work)
		fmt.Println("Merged: ", ws.Merged)
		fmt.Println("Status: ", status)
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}