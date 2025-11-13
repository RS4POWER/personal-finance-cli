package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/RS4POWER/personal-finance-cli/internal/db"
	"github.com/RS4POWER/personal-finance-cli/internal/repo"
)

var searchText string

func init() {
	searchCmd := &cobra.Command{
		Use:   "search",
		Short: "Search transactions by text in description",
		RunE:  runSearch,
	}

	searchCmd.Flags().StringVar(&searchText, "text", "", "Text to search in description")
	_ = searchCmd.MarkFlagRequired("text")

	rootCmd.AddCommand(searchCmd)
}

func runSearch(cmd *cobra.Command, args []string) error {
	database, err := db.Open("finance.db")
	if err != nil {
		return err
	}
	defer database.Close()

	r := repo.NewTransactionRepo(database)

	results, err := r.SearchByText(searchText)
	if err != nil {
		return err
	}

	if len(results) == 0 {
		fmt.Println("No transactions found.")
		return nil
	}

	for _, t := range results {
		fmt.Printf("[%d] %s | %-7s | %8.2f | %-10s | %s\n",
			t.ID,
			t.Date.Format("2006-01-02"),
			t.Type,
			t.Amount,
			t.Category,
			t.Description,
		)
	}

	return nil
}
