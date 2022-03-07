package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:   "app",
	Short: "skelaton-hexagonal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

// Execute is root function
func Execute() {
	rootCMD.AddCommand(restCMD)
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
