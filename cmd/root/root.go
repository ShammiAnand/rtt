package cmd

import (
	"github.com/shammianand/rtt/utils/logger"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "rtt",
		Short:   "Converts a directory to a text file",
		Long:    `rtt is a CLI tool that converts a directory to a text file with the same name as the directory.`,
		Example: `rtt /path/to/directory`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				logger.Log.Error("no args provided. use `.` to convert the current directory")
				return
			}

			logger.Log.Info("Hello from rtt", "args", args)
		},
	}
)

func Exectue() {
	rootCmd.Execute()
}
