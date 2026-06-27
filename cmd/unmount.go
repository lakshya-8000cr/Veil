package cmd

import (
	"fmt"

	"veil/internals/workspace"

	"github.com/spf13/cobra"
)

var unmountCmd = &cobra.Command{
	Use:   "unmount <name>",
	Short: "Unmount a Veil workspace",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		ws, err := workspace.Load(name)
		if err != nil {
			fmt.Println("failed to load workspace:", err)
			return
		}

		if err := ws.Unmount(); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println()
		fmt.Println("VEIL   Workspace unmounted")
		fmt.Println()
		fmt.Println("Name:  ", ws.Name)
		fmt.Println("Merged:", ws.Merged)
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(unmountCmd)
}