package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	importCmd := &cobra.Command{
		Use:   "import",
		Short: "Import transactions from a file (CSV/OFX)",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("import command not implemented yet (planned for CSV/OFX).")
			return nil
		},
	}

	rootCmd.AddCommand(importCmd)
}
