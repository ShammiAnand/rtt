package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/shammianand/rtt/pkg/chat"
	"github.com/shammianand/rtt/pkg/html"
	"github.com/shammianand/rtt/pkg/walker"
	"github.com/shammianand/rtt/utils/logger"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query [directory_path/url] [question]",
	Short: "ask questions about code or content",
	Long:  "requires exporting GROQ_API_KEY for chatting with an LLM",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			logger.Log.Error("Both source and question are required")
			return
		}

		apiKey := os.Getenv("GROQ_API_KEY")
		if apiKey == "" {
			logger.Log.Error("GROQ_API_KEY environment variable is required")
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		source := args[0]
		question := args[1]

		tempFile := fmt.Sprintf("%d.md", time.Now().Unix())

		if html.IsURL(source) {
			if err := html.ParseHTML(ctx, source, tempFile); err != nil {
				logger.Log.Error(err)
				return
			}
		} else {
			if err := walker.WalkAndExtract(source, tempFile); err != nil {
				logger.Log.Error(err)
				return
			}
		}

		if err := chat.ProcessQuery(ctx, tempFile, question, apiKey); err != nil {
			logger.Log.Error(err)
		}

		os.Remove(tempFile)

		fmt.Println()
		logger.Log.Info("File removed: ", tempFile)

	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}
