package cmd

import (
	"github.com/muhfaris/skeleton-hexagonal/transport/http/graphql"
	"github.com/spf13/cobra"
)

var graphqlCMD = &cobra.Command{
	Use: "graphql",
	Run: func(cmd *cobra.Command, args []string) {
		graphql.NewGraphql(5555).Serve()
	},
}
