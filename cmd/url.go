package cmd

import (
	"github.com/shammianand/rtt/utils/logger"
	"github.com/spf13/cobra"
)

var (
	urlCommand = &cobra.Command{
		Use:     "url",
		Short:   "converts the provided url to a markdown file",
		Long:    "converts the provided url to a markdown file",
		Example: "rtt url https://www.google.com",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Log.Info("args: ", args)
		},
	}
)

func init() {
	rootCmd.AddCommand(urlCommand)
}
