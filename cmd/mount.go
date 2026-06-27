package cmd

import (
	"fmt"

	"veil/internals/overlay"
	"veil/internals/workspace"

	"github.com/spf13/cobra"
)

var mountCmd = &cobra.Command{
	Use:   "mount <name>",
	Short: "Mount a Veil workspace using OverlayFS",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		ws, err := workspace.Load(name)
		if err != nil {
			fmt.Println("failed to load workspace:", err)
			return
		}

		if err := overlay.Mount(ws.Project, ws.Upper, ws.Work, ws.Merged); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println()
		fmt.Println("VEIL   Workspace mounted")
		fmt.Println()
		fmt.Println("Name:   ", ws.Name)
		fmt.Println("Project:", ws.Project)
		fmt.Println("Merged: ", ws.Merged)
		fmt.Println()
		fmt.Println("Open:")
		fmt.Println("  code", ws.Merged)
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(mountCmd)
}