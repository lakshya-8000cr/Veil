package cmd

import (
	"fmt"
	"time"

	"veil/internals/workspace"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect <name>",
	Short: "Inspect a Veil workspace",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[14], 80*time.Millisecond)
		s.Suffix = "  loading workspace"
		s.Start()

		ws, err := workspace.Load(args[0])

		time.Sleep(2 * time.Second)
		s.Stop()

		if err != nil {
			fmt.Printf("  %s  loading workspace\n", color.RedString("✖"))
			fmt.Println("failed to load workspace:", err)
			return
		}

		status := "unmounted"
		if ws.IsMounted() {
			status = "mounted"
		}

		green := color.New(color.FgGreen).SprintFunc()
		dim := color.New(color.FgHiBlack).SprintFunc()
		white := color.New(color.FgWhite).SprintFunc()
		cyan := color.New(color.FgCyan).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()

		fmt.Println()
		color.New(color.FgWhite, color.Bold).Print("veil")
		color.New(color.FgHiBlack).Print("  ›  ")
		color.New(color.FgWhite).Printf("inspect %s\n", ws.Name)
		fmt.Println()

		fmt.Printf("  %s  name      %s %s\n", green("✔"), dim("→"), white(ws.Name))
		fmt.Printf("  %s  project   %s %s\n", green("✔"), dim("→"), white(ws.Project))
		fmt.Printf("  %s  upper     %s %s\n", green("✔"), dim("→"), dim(ws.Upper))
		fmt.Printf("  %s  work      %s %s\n", green("✔"), dim("→"), dim(ws.Work))
		fmt.Printf("  %s  merged    %s %s\n", green("✔"), dim("→"), white(ws.Merged))

		if status == "mounted" {
			fmt.Printf("  %s  status    %s %s\n", green("✔"), dim("→"), cyan(status))
		} else {
			fmt.Printf("  %s  status    %s %s\n", green("✔"), dim("→"), yellow(status))
		}

		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}