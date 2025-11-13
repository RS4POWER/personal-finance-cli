package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	budgetCmd := &cobra.Command{
		Use:   "budget",
		Short: "Manage budgets per category",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("budget command not implemented yet (planned for limits & alerts).")
			return nil
		},
	}

	rootCmd.AddCommand(budgetCmd)
}
