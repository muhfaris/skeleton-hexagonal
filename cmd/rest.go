package cmd

import (
	"github.com/muhfaris/skeleton-hexagonal/transport/http/rest"
	"github.com/spf13/cobra"
)

var restCMD = &cobra.Command{
	Use: "api",
	Run: func(cmd *cobra.Command, args []string) {
		rest.NewRest(5555).Serve()
	},
}
