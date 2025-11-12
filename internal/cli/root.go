package cli

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "pfcli",
	Short: "Personal Finance CLI Manager",
	Long:  "Track income and expenses from your terminal.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// aici vom baga subcomenzile: add, import, report, search, budget, tui
}
