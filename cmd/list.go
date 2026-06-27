package cmd

import (
	"fmt"

	"veil/internals/workspace"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Veil workspaces",

	Run: func(cmd *cobra.Command, args []string) {
		workspaces, err := workspace.List()
		if err != nil {
			fmt.Println("failed to list workspaces:", err)
			return
		}

		fmt.Println()
		fmt.Println("VEIL   Workspaces")
		fmt.Println()

		if len(workspaces) == 0 {
			fmt.Println("No workspaces found.")
			fmt.Println()
			return
		}

		fmt.Printf(" %-20s %-60s %s\n", "NAME", "PROJECT", "MERGED")
		fmt.Println()

		for _, ws := range workspaces {
			fmt.Printf(
				" %-20s %-60s %s\n",
				ws.Name,
				ws.Project,
				ws.Merged,
			)
		}

		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}