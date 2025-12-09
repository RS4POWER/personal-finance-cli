package cli

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/RS4POWER/personal-finance-cli/internal/db"
	"github.com/RS4POWER/personal-finance-cli/internal/domain"
	"github.com/RS4POWER/personal-finance-cli/internal/repo"
)

var importFile string

func init() {
	importCmd := &cobra.Command{
		Use:   "import",
		Short: "Import transactions from a CSV file",
		Long: `Import transactions from a CSV file into the local SQLite database.

Expected CSV header (case-insensitive):
  date, description, amount, category, type

- date:        YYYY-MM-DD (if empty, today is used)
- amount:      numeric
- type:        income|expense (defaults to expense if empty)
- category:    free text (optional)`,
		RunE: runImport,
	}

	importCmd.Flags().StringVarP(&importFile, "file", "f", "", "Path to CSV file")
	_ = importCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(importCmd)
}

func runImport(cmd *cobra.Command, args []string) error {
	f, err := os.Open(importFile)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.TrimLeadingSpace = true

	// Read header
	header, err := reader.Read()
	if err != nil {
		return fmt.Errorf("read header: %w", err)
	}

	// Map column name -> index (lowercased)
	colIndex := map[string]int{}
	for i, h := range header {
		name := strings.ToLower(strings.TrimSpace(h))
		colIndex[name] = i
	}

	// At minimum we need: description, amount
	descIdx, ok := colIndex["description"]
	if !ok {
		return fmt.Errorf("CSV must contain a 'description' column")
	}
	amountIdx, ok := colIndex["amount"]
	if !ok {
		return fmt.Errorf("CSV must contain an 'amount' column")
	}

	// Optional columns
	dateIdx, hasDate := colIndex["date"]
	categoryIdx, hasCategory := colIndex["category"]
	typeIdx, hasType := colIndex["type"]

	database, err := db.Open("finance.db")
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}
	defer database.Close()

	r := repo.NewTransactionRepo(database)
	ruleRepo := repo.NewRuleRepo(database)

	var imported int
	var failed int

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("read row: %w", err)
		}

		// Description
		description := strings.TrimSpace(record[descIdx])
		if description == "" {
			failed++
			continue
		}

		// Amount
		amountStr := strings.TrimSpace(record[amountIdx])
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			fmt.Printf("Skipping row (invalid amount '%s'): %v\n", amountStr, err)
			failed++
			continue
		}

		// Date
		var date time.Time
		if hasDate {
			dateStr := strings.TrimSpace(record[dateIdx])
			if dateStr == "" {
				date = time.Now()
			} else {
				d, err := time.Parse("2006-01-02", dateStr)
				if err != nil {
					fmt.Printf("Skipping row (invalid date '%s'): %v\n", dateStr, err)
					failed++
					continue
				}
				date = d
			}
		} else {
			date = time.Now()
		}

		// Category
		category := ""
		if hasCategory {
			category = strings.TrimSpace(record[categoryIdx])
		}
		// If category is empty, try rules
		if category == "" {
			if cat, err := ruleRepo.FindCategory(description); err == nil && cat != "" {
				category = cat
			}
		}

		// Type
		tType := domain.TransactionTypeExpense
		if hasType {
			typeStr := strings.ToLower(strings.TrimSpace(record[typeIdx]))
			if typeStr == "income" {
				tType = domain.TransactionTypeIncome
			}
		}

		tx := &domain.Transaction{
			Date:        date,
			Description: description,
			Amount:      amount,
			Category:    category,
			Type:        tType,
		}

		if err := r.Insert(tx); err != nil {
			fmt.Printf("Failed to insert transaction '%s': %v\n", description, err)
			failed++
			continue
		}

		imported++
	}

	fmt.Printf("Import finished. Imported: %d, Skipped: %d\n", imported, failed)
	return nil
}
