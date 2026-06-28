package cmd

import (
	"fmt"

	"veil/internals/workspace"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy <name>",
	Short: "Destroy a Veil workspace",
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
		color.New(color.FgWhite).Printf("destroying %s\n", name)
		fmt.Println()

		green := color.New(color.FgGreen).SprintFunc()
		dim := color.New(color.FgHiBlack).SprintFunc()
		white := color.New(color.FgWhite).SprintFunc()

		if err := ws.Destroy(); err != nil {
			fmt.Printf("  %s  tearing down workspace\n", color.RedString("✖"))
			fmt.Println("failed to destroy workspace:", err)
			return
		}

		fmt.Printf("  %s  workspace destroyed\n", green("✔"))
		fmt.Printf("  %s  name   %s %s\n", green("✔"), dim("→"), white(ws.Name))
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)
}