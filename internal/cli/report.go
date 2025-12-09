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

		// find maximum total for scaling the ASCII bar
		max := 0.0
		for _, c := range cats {
			if c.Total > max {
				max = c.Total
			}
		}

		const barMax = 25 // length of the bar

		// ANSI colors for the bar
		red := "\033[31m"
		green := "\033[32m"
		yellow := "\033[33m"
		reset := "\033[0m"

		for _, c := range cats {
			category := c.Category
			if category == "" {
				category = "(uncategorized)"
			}

			// scale bar length
			barLen := int((c.Total / max) * barMax)
			if barLen < 2 {
				barLen = 2 // always at least visible
			}

			bar := ""
			for i := 0; i < barLen; i++ {
				bar += "█"
			}

			// default: no budget -> yellow bar, empty status
			statusText := ""
			color := yellow

			if limit, ok := budgetMap[category]; ok {
				// there's a budget for this category
				if c.Total > limit {
					statusText = fmt.Sprintf("/ %.2f (OVER)", limit)
					color = red
				} else {
					statusText = fmt.Sprintf("/ %.2f (OK)", limit)
					color = green
				}
			}

			// left side: category, total, status text (no color, fixed width)
			left := fmt.Sprintf("  %-12s %-8.2f %-20s", category, c.Total, statusText)

			// colored bar
			coloredBar := color + bar + reset

			fmt.Printf("%s %s\n", left, coloredBar)
		}
	}

	return nil
}
