package cmd

import (
	"fmt"

	"veil/internals/workspace"

	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch <name>",
	Short: "Watch workspace file changes using inotify",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		ws, err := workspace.Load(args[0])
		if err != nil {
			fmt.Println("failed to load workspace:", err)
			return
		}

		if !ws.IsMounted() {
			fmt.Println("workspace is not mounted. Run:")
			fmt.Println("  veil mount", ws.Name)
			return
		}

		if err := ws.Watch(); err != nil {
			fmt.Println("watch failed:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}