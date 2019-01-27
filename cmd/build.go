package cmd

import (
	"github.com/joway/imagic/pkg/util"
	"github.com/spf13/cobra"
	"path/filepath"
)

func init() {
	rootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
	Use: "build",
	Run: func(cmd *cobra.Command, args []string) {
		patterns := args
		// get all files
		files := make([]string, 0)
		for _, pattern := range patterns {
			_files, err := filepath.Glob(pattern)
			if err != nil {
				util.LogFatal(err)
			}
			files = append(files, _files...)
		}

	},
}
