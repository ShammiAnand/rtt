package cmd

import (
	"github.com/shammianand/rtt/pkg/walker"
	"github.com/shammianand/rtt/utils/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func GetCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		logger.Log.Error(err)
	}
	logger.Log.Info("using current dir: ", dir)
	return dir
}

var (
	outputFile string
	logLevel   string
	rootCmd    = &cobra.Command{
		Use:     "rtt",
		Short:   "Converts a directory to a `md` file",
		Long:    `rtt is a CLI tool that converts a directory to a markdown file with the same name as the directory.`,
		Example: `rtt /path/to/directory`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) > 1 {
				logger.Log.Error("only one dir at a time is allowed")
				return
			}

			var walkPath string
			if len(args) == 0 {
				walkPath = GetCurrentDir()
			} else {
				walkPath = args[0]
				if walkPath == "." {
					walkPath = GetCurrentDir()
				}
			}

			if err := walker.WalkAndExtract(walkPath, outputFile); err != nil {
				logger.Log.Error(err)
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringP("author", "a", "Shammi Anand", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "rtt.md", "output file name")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log", "l", "INFO", "verbosity of the logger")

	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.SetDefault("author", "Shammi Anand shammianand101@gmail.com")
	viper.SetDefault("license", "apache")

	logger.InitLogger(logLevel)
}

func Exectue() {
	rootCmd.Execute()
}
