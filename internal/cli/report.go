package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/RS4POWER/personal-finance-cli/internal/db"
	"github.com/RS4POWER/personal-finance-cli/internal/repo"
)

func init() {
	reportCmd := &cobra.Command{
		Use:   "report",
		Short: "Show a simple income/expense summary",
		RunE:  runReport,
	}

	rootCmd.AddCommand(reportCmd)
}

func runReport(cmd *cobra.Command, args []string) error {
	database, err := db.Open("finance.db")
	if err != nil {
		return err
	}
	defer database.Close()

	r := repo.NewTransactionRepo(database)

	income, expense, err := r.Totals()
	if err != nil {
		return err
	}

	balance := income - expense

	fmt.Printf("Income : %.2f\n", income)
	fmt.Printf("Expense: %.2f\n", expense)
	fmt.Printf("Balance: %.2f\n", balance)

	return nil
}
