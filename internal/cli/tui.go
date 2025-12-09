package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/RS4POWER/personal-finance-cli/internal/db"
	"github.com/RS4POWER/personal-finance-cli/internal/repo"
)

func init() {
	tuiCmd := &cobra.Command{
		Use:   "tui",
		Short: "Open a simple TUI to view recent transactions",
		RunE:  runTUI,
	}

	rootCmd.AddCommand(tuiCmd)
}

func runTUI(cmd *cobra.Command, args []string) error {
	database, err := db.Open("finance.db")
	if err != nil {
		return err
	}
	defer database.Close()

	r := repo.NewTransactionRepo(database)

	txs, err := r.LastN(10)
	if err != nil {
		return err
	}

	fmt.Println("=== Recent Transactions ===")
	fmt.Println("ID   DATE         TYPE       AMOUNT     CATEGORY      DESCRIPTION")
	fmt.Println("--------------------------------------------------------------------------")

	for _, t := range txs {
		fmt.Printf("%-4d %-12s %-10s %-10.2f %-12s %s\n",
			t.ID,
			t.Date.Format("2006-01-02"),
			string(t.Type),
			t.Amount,
			t.Category,
			t.Description,
		)
	}

	fmt.Println("\nPress ENTER to exit...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadBytes('\n')

	return nil
}
