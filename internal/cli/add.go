package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/RS4POWER/personal-finance-cli/internal/db"
	"github.com/RS4POWER/personal-finance-cli/internal/domain"
	"github.com/RS4POWER/personal-finance-cli/internal/repo"
)

var (
	addAmount      float64
	addDescription string
	addCategory    string
	addType        string
	addDate        string
)

func init() {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a transaction manually",
		RunE:  runAdd,
	}

	addCmd.Flags().Float64VarP(&addAmount, "amount", "a", 0, "Amount")
	addCmd.Flags().StringVarP(&addDescription, "description", "d", "", "Description")
	addCmd.Flags().StringVarP(&addCategory, "category", "c", "", "Category")
	addCmd.Flags().StringVarP(&addType, "type", "t", "expense", "Type: income|expense")
	addCmd.Flags().StringVar(&addDate, "date", "", "Date (YYYY-MM-DD), defaults to today")

	addCmd.MarkFlagRequired("amount")
	addCmd.MarkFlagRequired("description")

	rootCmd.AddCommand(addCmd)
}

func runAdd(cmd *cobra.Command, args []string) error {
	if addDate == "" {
		addDate = time.Now().Format("2006-01-02")
	}
	parsedDate, err := time.Parse("2006-01-02", addDate)
	if err != nil {
		return err
	}

	tType := domain.TransactionTypeExpense
	if addType == "income" {
		tType = domain.TransactionTypeIncome
	}

	database, err := db.Open("finance.db")
	if err != nil {
		return err
	}
	defer database.Close()

	r := repo.NewTransactionRepo(database)

	tx := &domain.Transaction{
		Date:        parsedDate,
		Description: addDescription,
		Amount:      addAmount,
		Category:    addCategory,
		Type:        tType,
	}

	if err := r.Insert(tx); err != nil {
		return err
	}

	fmt.Println("Transaction added with ID:", tx.ID)
	return nil
}
