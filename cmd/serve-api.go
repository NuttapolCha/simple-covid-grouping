package cmd

import "github.com/spf13/cobra"

var serveApiCommand = &cobra.Command{
	Use:   "serve-api",
	Short: "Start API Server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
