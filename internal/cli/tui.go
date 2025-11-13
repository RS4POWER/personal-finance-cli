package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	tuiCmd := &cobra.Command{
		Use:   "tui",
		Short: "Start the interactive TUI",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TUI not implemented yet (planned interactive interface).")
			return nil
		},
	}

	rootCmd.AddCommand(tuiCmd)
}
