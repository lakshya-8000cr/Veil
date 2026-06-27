package cmd

import (
	"fmt"

	"veil/internals/workspace"

	"github.com/spf13/cobra"
)
  
var spawnCmd = &cobra.Command{
Use:   "spawn <name> [project-path]",
Short: "Create a disposable workspace",
Args:  cobra.RangeArgs(1, 2),

	Run: func(cmd *cobra.Command, args []string) {
	name := args[0]

	projectPath := "."
	if len(args) == 2 {
		projectPath = args[1]
	}

	ws, err := workspace.New(name, projectPath)
	if err != nil {
		fmt.Println("failed to create workspace:", err)
		return
	}

	if err := ws.Create(); err != nil {
		fmt.Println("failed to initialize workspace:", err)
		return
	}

	fmt.Println()
	fmt.Println("VEIL   Workspace created")
	fmt.Println()
	fmt.Println("Name:    ", ws.Name)
	fmt.Println("Project: ", ws.Project)
	fmt.Println("Merged:  ", ws.Merged)
	fmt.Println()
	fmt.Println("Next:")
	fmt.Println("  veil mount", ws.Name)
	fmt.Println("  code", ws.Merged)
	fmt.Println()
},
}

func init() {
	rootCmd.AddCommand(spawnCmd)
}