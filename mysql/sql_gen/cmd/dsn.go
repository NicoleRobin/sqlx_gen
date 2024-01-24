package cmd

import "github.com/spf13/cobra"

var dsn = &cobra.Command{
	Use:   "dsn",
	Short: "generate code from dsn",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(dsn)
}
