package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/RS4POWER/personal-finance-cli/internal/db"
	"github.com/RS4POWER/personal-finance-cli/internal/domain"
	"github.com/RS4POWER/personal-finance-cli/internal/repo"
)

var reportByCategory bool

func init() {
	reportCmd := &cobra.Command{
		Use:   "report",
		Short: "Show a summary of income/expense and optional category breakdown",
		RunE:  runReport,
	}

	reportCmd.Flags().BoolVar(&reportByCategory, "by-category", false, "Show expense totals by category")

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

	if reportByCategory {
		fmt.Println()
		fmt.Println("Expense by category:")

		cats, err := r.TotalsByCategory(domain.TransactionTypeExpense)
		if err != nil {
			return err
		}

		if len(cats) == 0 {
			fmt.Println("  (no expenses found)")
			return nil
		}

		// Load budgets
		budgetRepo := repo.NewBudgetRepo(database)
		budgets, err := budgetRepo.List()
		if err != nil {
			return err
		}
		budgetMap := make(map[string]float64)
		for _, b := range budgets {
			budgetMap[b.Category] = b.Limit
		}

		for _, c := range cats {
			category := c.Category
			if category == "" {
				category = "(uncategorized)"
			}

			limit, hasBudget := budgetMap[category]
			if hasBudget {
				status := "OK"
				if c.Total > limit {
					status = "OVER"
				}
				fmt.Printf("  %-15s %.2f / %.2f (%s)\n", category, c.Total, limit, status)
			} else {
				fmt.Printf("  %-15s %.2f\n", category, c.Total)
			}
		}
	}

	return nil
}
