package cmd

import (
	"fmt"

	"veil/internals/overlay"
	"veil/internals/workspace"

	"github.com/fatih/color"
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

		fmt.Println()
		color.New(color.FgWhite, color.Bold).Print("veil")
		color.New(color.FgHiBlack).Print("  ›  ")
		color.New(color.FgWhite).Printf("mounting %s\n", name)
		fmt.Println()

		green := color.New(color.FgGreen).SprintFunc()
		dim := color.New(color.FgHiBlack).SprintFunc()
		white := color.New(color.FgWhite).SprintFunc()
		cyan := color.New(color.FgCyan).SprintFunc()

		if err := overlay.Mount(ws.Project, ws.Upper, ws.Work, ws.Merged); err != nil {
			fmt.Printf("  %s  overlayfs mounting\n", color.RedString("✖"))
			fmt.Println(err)
			return
		}

		fmt.Printf("  %s  overlayfs mounted\n", green("✔"))
		fmt.Printf("  %s  project linked   %s %s\n", green("✔"), dim("→"), white(ws.Project))
		fmt.Printf("  %s  merged ready     %s %s\n", green("✔"), dim("→"), white(ws.Merged))
		fmt.Println()
		fmt.Printf("  %s    %s\n", dim("open   "), cyan("code "+ws.Merged))
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(mountCmd)
}