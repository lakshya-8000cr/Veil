package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"veil/internals/workspace"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var spawnCmd = &cobra.Command{
	Use:   "spawn [name] [project-path]",
	Short: "Create a disposable workspace",
	Args:  cobra.MaximumNArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		projectPath := "."
		name := ""

		if len(args) == 0 {
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Println("failed to get current directory:", err)
				return
			}

			name = filepath.Base(cwd)
		}

		if len(args) == 1 {
			name = args[0]
		}

		if len(args) == 2 {
			name = args[0]
			projectPath = args[1]
		}

		existing, err := workspace.FindByProject(projectPath)
		if err != nil {
			fmt.Println("failed to check existing workspaces:", err)
			return
		}

		if existing != nil {
			fmt.Println()
			color.New(color.FgWhite, color.Bold).Print("veil")
			color.New(color.FgHiBlack).Print("  ›  ")
			color.New(color.FgYellow).Printf("workspace already exists  %s\n", existing.Name)
			fmt.Println()
			color.New(color.FgHiBlack).Print("  name     ")
			color.New(color.FgWhite).Println(existing.Name)
			color.New(color.FgHiBlack).Print("  project  ")
			color.New(color.FgWhite).Println(existing.Project)
			fmt.Println()
			color.New(color.FgHiBlack).Print("  mount    ")
			color.New(color.FgCyan).Printf("veil mount %s\n", existing.Name)
			color.New(color.FgHiBlack).Print("  inspect  ")
			color.New(color.FgCyan).Printf("veil inspect %s\n", existing.Name)
			fmt.Println()
			return
		}

		ws, err := workspace.New(name, projectPath)
		if err != nil {
			fmt.Println("failed to create workspace:", err)
			return
		}

		fmt.Println()
		color.New(color.FgWhite, color.Bold).Print("veil")
		color.New(color.FgHiBlack).Print("  ›  ")
		color.New(color.FgWhite).Printf("spawning %s\n", name)
		fmt.Println()

		green := color.New(color.FgGreen).SprintFunc()
		dim := color.New(color.FgHiBlack).SprintFunc()
		white := color.New(color.FgWhite).SprintFunc()
		cyan := color.New(color.FgCyan).SprintFunc()

		s := spinner.New(spinner.CharSets[14], 80*time.Millisecond)
		s.Suffix = "  creating workspace"
		s.Start()

		createErr := ws.Create()

		time.Sleep(2 * time.Second)
		s.Stop()

		if createErr != nil {
			fmt.Printf("  %s  creating workspace\n", color.RedString("✖"))
			fmt.Println("failed to initialize workspace:", createErr)
			return
		}

		fmt.Printf("  %s  overlay created\n", green("✔"))
		fmt.Printf("  %s  project linked   %s %s\n", green("✔"), dim("→"), white(ws.Project))
		fmt.Printf("  %s  workspace ready  %s %s\n", green("✔"), dim("→"), white(ws.Merged))
		fmt.Println()
		fmt.Printf("  %s    %s\n", dim("mount  "), cyan("veil mount "+ws.Name))
		fmt.Printf("  %s  %s\n", dim("inspect"), cyan("veil inspect "+ws.Name))
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(spawnCmd)
}