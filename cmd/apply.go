package cmd

import (
	"fmt"
	"time"

	"veil/internals/workspace"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply <name>",
	Short: "Apply workspace changes back to the original project",
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
		color.New(color.FgWhite).Printf("applying %s\n", name)
		fmt.Println()

		green := color.New(color.FgGreen).SprintFunc()
		dim := color.New(color.FgHiBlack).SprintFunc()
		white := color.New(color.FgWhite).SprintFunc()

		s := spinner.New(spinner.CharSets[14], 80*time.Millisecond)
		s.Suffix = "  copying changes to project"
		s.Start()

		applyErr := ws.Apply()

		time.Sleep(2 * time.Second)
		s.Stop()

		if applyErr != nil {
			fmt.Printf("  %s  copying changes to project\n", color.RedString("✖"))
			fmt.Println("failed to apply changes:", applyErr)
			return
		}

		fmt.Printf("  %s  changes applied\n", green("✔"))
		fmt.Printf("  %s  workspace   %s %s\n", green("✔"), dim("→"), white(ws.Name))
		fmt.Printf("  %s  project     %s %s\n", green("✔"), dim("→"), white(ws.Project))
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}