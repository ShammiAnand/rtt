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
	outputFile     string
	logLevel       string
	skipCacheFiles bool
	optimizedText  bool
	rootCmd        = &cobra.Command{
		Use:     "rtt [path]",
		Short:   "Converts a directory to a `md` or optimized `txt` file",
		Long:    `rtt is a CLI tool that converts a directory to a markdown file or optimized text file with the same name as the directory.`,
		Example: `rtt /path/to/directory
rtt . -o output.txt --optimized-text
rtt --skip-cache -o optimized.txt --optimized-text`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var walkPath string
			if len(args) == 0 {
				walkPath = GetCurrentDir()
			} else {
				walkPath = args[0]
				if walkPath == "." {
					walkPath = GetCurrentDir()
				}
			}
			
			// Create file filter based on flags
			filter := walker.FileFilter{
				SkipCacheFiles: skipCacheFiles,
				OptimizedText:  optimizedText,
			}
			
			return walker.WalkAndExtract(walkPath, outputFile, filter)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringP("author", "a", "Shammi Anand", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "rtt.md", "output file name")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log", "l", "INFO", "verbosity of the logger")
	rootCmd.PersistentFlags().BoolVarP(&skipCacheFiles, "skip-cache", "s", false, "skip cache files and non-code/text files")
	rootCmd.PersistentFlags().BoolVarP(&optimizedText, "optimized-text", "t", false, "create optimized text file with minimal formatting to save tokens")

	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.SetDefault("author", "Shammi Anand shammianand101@gmail.com")
	viper.SetDefault("license", "apache")

	logger.InitLogger(logLevel)
}

func Execute() {
	rootCmd.Execute()
}
