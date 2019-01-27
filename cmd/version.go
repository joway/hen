package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = "v0.0.1"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hen %s\n", version)
	},
}
