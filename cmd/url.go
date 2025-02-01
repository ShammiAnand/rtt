package cmd

import (
	"context"
	"time"

	"github.com/shammianand/rtt/pkg/html"
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
			if len(args) == 0 {
				logger.Log.Error("url is required")
				return
			}

			// we don't want to wait for more than 10 seconds
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			err := html.ParseHTML(ctx, args[0])
			if err != nil {
				logger.Log.Error("Error parsing HTML: ", err)
				return
			}

		},
	}
)

func init() {
	rootCmd.AddCommand(urlCommand)
}
