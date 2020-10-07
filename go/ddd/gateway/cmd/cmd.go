package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	dbPool *sql.DB
)

var rootCmd = &cobra.Command{
	Use: "adsrobot",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
			logrus.SetLevel(logrus.DebugLevel)
		}
	},
}

func init() {
	cobra.OnInitialize(splash, initconfig, initDatabase, initCache)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
}

// Execute is root function
func Execute() {
	rootCmd.AddCommand(serveCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
