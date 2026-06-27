package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"veil/internals/workspace"

	"github.com/spf13/cobra"
)

var spawnCmd = &cobra.Command{
	Use:   "spawn [name] [project-path]",
	Short: "Create a disposable workspace",
	Args:  cobra.MaximumNArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		projectPath := "."
		name := ""

		if len(args) == 0 {
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Println("failed to get current directory:", err)
				return
			}

			name = filepath.Base(cwd)
		}

		if len(args) == 1 {
			name = args[0]
		}

		if len(args) == 2 {
			name = args[0]
			projectPath = args[1]
		}

		existing, err := workspace.FindByProject(projectPath)
		if err != nil {
			fmt.Println("failed to check existing workspaces:", err)
			return
		}

		if existing != nil {
			fmt.Println()
			fmt.Println("VEIL   Workspace already exists")
			fmt.Println()
			fmt.Println("Name:   ", existing.Name)
			fmt.Println("Project:", existing.Project)
			fmt.Println()
			fmt.Println("Use:")
			fmt.Println("  veil mount", existing.Name)
			fmt.Println("  veil inspect", existing.Name)
			fmt.Println()
			return
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
