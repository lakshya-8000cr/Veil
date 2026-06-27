package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{  // root command
	use:  "veil",
	Short:  "Disposable development workspaces using OverlayFS",
}

func Execute(){
     if err := rootCmd.Execute(); err!= nil {  // if any unkonown is there then print the error
		fmt.Println(err)
		os.Exit(1)
	 }
}