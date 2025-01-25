package cmd

import (
	"github.com/shammianand/rtt/pkg/walker"
	"github.com/shammianand/rtt/utils/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	outputFile string
	rootCmd    = &cobra.Command{
		Use:     "rtt",
		Short:   "Converts a directory to a text file",
		Long:    `rtt is a CLI tool that converts a directory to a text file with the same name as the directory.`,
		Example: `rtt /path/to/directory`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				logger.Log.Error("no args provided. use `.` to convert the current directory")
				return
			}

			if len(args) > 1 {
				logger.Log.Error("only one argument is allowed")
				return
			}

			// logger.Log.Info("author: ", viper.GetString("author"))
			// logger.Log.Info("output: ", outputFile)
			// logger.Log.Info("args: ", args)

			if err := walker.WalkAndExtract(args[0], outputFile); err != nil {
				logger.Log.Error(err)
			}
		},
	}
)

func Exectue() {
	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("author", "a", "Shammi Anand", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "rtt-<dir_name>.txt", "output file name")

	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.SetDefault("author", "Shammi Anand shammianand101@gmail.com")
	viper.SetDefault("license", "apache")

}
