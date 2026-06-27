package cmd

import (
	"fmt"
	"time"

	"veil/internals/workspace"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
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

		fmt.Println()
		color.New(color.FgWhite, color.Bold).Print("veil")
		color.New(color.FgHiBlack).Print("  ›  ")
		color.New(color.FgWhite).Printf("unmounting %s\n", name)
		fmt.Println()

		green := color.New(color.FgGreen).SprintFunc()
		dim := color.New(color.FgHiBlack).SprintFunc()
		white := color.New(color.FgWhite).SprintFunc()

		s := spinner.New(spinner.CharSets[14], 80*time.Millisecond)
		s.Suffix = "  detaching overlayfs"
		s.Start()

		unmountErr := ws.Unmount()

		time.Sleep(2 * time.Second)
		s.Stop()

		if unmountErr != nil {
			fmt.Printf("  %s  detaching overlayfs\n", color.RedString("✖"))
			fmt.Println(unmountErr)
			return
		}

		fmt.Printf("  %s  overlayfs detached\n", green("✔"))
		fmt.Printf("  %s  name     %s %s\n", green("✔"), dim("→"), white(ws.Name))
		fmt.Printf("  %s  merged   %s %s\n", green("✔"), dim("→"), white(ws.Merged))
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(unmountCmd)
}