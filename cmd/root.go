package cmd

import (
	"github.com/joway/imagic/pkg/util"
	"github.com/spf13/cobra"
)

func init() {
}

var rootCmd = &cobra.Command{
	Use:   "hen",
	Short: "A static site generator.",
	//Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

// Execute root cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		util.LogFatal(err)
	}
}
