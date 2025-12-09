package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/RS4POWER/personal-finance-cli/internal/db"
	"github.com/RS4POWER/personal-finance-cli/internal/domain"
	"github.com/RS4POWER/personal-finance-cli/internal/repo"
)

var (
	budgetCategory string
	budgetLimit    float64
)

func init() {
	budgetCmd := &cobra.Command{
		Use:   "budget",
		Short: "Manage budgets per category",
	}

	addBudgetCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a budget for a category",
		RunE:  runAddBudget,
	}
	addBudgetCmd.Flags().StringVar(&budgetCategory, "category", "", "Category name")
	addBudgetCmd.Flags().Float64Var(&budgetLimit, "limit", 0, "Limit amount")
	_ = addBudgetCmd.MarkFlagRequired("category")
	_ = addBudgetCmd.MarkFlagRequired("limit")

	listBudgetCmd := &cobra.Command{
		Use:   "list",
		Short: "List all budgets",
		RunE:  runListBudgets,
	}

	budgetCmd.AddCommand(addBudgetCmd)
	budgetCmd.AddCommand(listBudgetCmd)

	rootCmd.AddCommand(budgetCmd)
}

func runAddBudget(cmd *cobra.Command, args []string) error {
	database, err := db.Open("finance.db")
	if err != nil {
		return err
	}
	defer database.Close()

	r := repo.NewBudgetRepo(database)

	b := &domain.Budget{
		Category: budgetCategory,
		Limit:    budgetLimit,
	}

	if err := r.Insert(b); err != nil {
		return err
	}

	fmt.Printf("Budget added for category '%s' with limit %.2f\n", b.Category, b.Limit)
	return nil
}

func runListBudgets(cmd *cobra.Command, args []string) error {
	database, err := db.Open("finance.db")
	if err != nil {
		return err
	}
	defer database.Close()

	r := repo.NewBudgetRepo(database)
	budgets, err := r.List()
	if err != nil {
		return err
	}

	if len(budgets) == 0 {
		fmt.Println("No budgets defined.")
		return nil
	}

	for _, b := range budgets {
		fmt.Printf("[%d] %-15s %.2f\n", b.ID, b.Category, b.Limit)
	}

	return nil
}
