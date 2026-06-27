package cmd

import (
	"fmt"
	"time"

	"veil/internals/workspace"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Veil workspaces",

	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[14], 80*time.Millisecond)
		s.Suffix = "  loading workspaces"
		s.Start()

		workspaces, err := workspace.List()

		time.Sleep(2 * time.Second)
		s.Stop()

		if err != nil {
			fmt.Printf("  %s  loading workspaces\n", color.RedString("✖"))
			fmt.Println("failed to list workspaces:", err)
			return
		}

		dim := color.New(color.FgHiBlack).SprintFunc()
		white := color.New(color.FgWhite, color.Bold).SprintFunc()
		cyan := color.New(color.FgCyan).SprintFunc()

		fmt.Println()
		color.New(color.FgWhite, color.Bold).Print("veil")
		color.New(color.FgHiBlack).Print("  ›  ")
		color.New(color.FgWhite).Println("workspaces")
		fmt.Println()

		if len(workspaces) == 0 {
			fmt.Printf("  %s  no workspaces found\n", dim("·"))
			fmt.Println()
			return
		}

		fmt.Printf("  %s   %-20s %-60s %s\n", dim("·"), dim("NAME"), dim("PROJECT"), dim("MERGED"))
		fmt.Println()

		for _, ws := range workspaces {
			fmt.Printf(
				"  %s   %-20s %-60s %s\n",
				color.GreenString("✔"),
				white(ws.Name),
				cyan(ws.Project),
				dim(ws.Merged),
			)
		}

		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}