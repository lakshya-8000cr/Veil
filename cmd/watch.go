package cmd

import (
	"fmt"

	"veil/internals/workspace"

	"github.com/fatih/color"
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
			fmt.Println()
			color.New(color.FgWhite, color.Bold).Print("veil")
			color.New(color.FgHiBlack).Print("  ›  ")
			color.New(color.FgYellow).Println("workspace not mounted")
			fmt.Println()
			color.New(color.FgHiBlack).Print("  mount   ")
			color.New(color.FgCyan).Printf("veil mount %s\n", ws.Name)
			fmt.Println()
			return
		}

		fmt.Println()
		color.New(color.FgWhite, color.Bold).Print("veil")
		color.New(color.FgHiBlack).Print("  ›  ")
		color.New(color.FgWhite).Printf("watching %s\n", ws.Name)
		fmt.Println()
		color.New(color.FgHiBlack).Printf("  %s  %s\n", "·", "listening for file changes  ctrl+c to stop")
		fmt.Println()

		if err := ws.Watch(); err != nil {
			fmt.Printf("  %s  watch failed: %s\n", color.RedString("✖"), err)
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}