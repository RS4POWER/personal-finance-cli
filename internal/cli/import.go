package cli

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/RS4POWER/personal-finance-cli/internal/db"
	"github.com/RS4POWER/personal-finance-cli/internal/domain"
	"github.com/RS4POWER/personal-finance-cli/internal/repo"
	"github.com/RS4POWER/personal-finance-cli/internal/service"
)

var importFile string

func init() {
	importCmd := &cobra.Command{
		Use:   "import",
		Short: "Import transactions from a CSV or OFX file",
		Long: `Import transactions from a CSV or OFX file into the local SQLite database.

CSV expected header (case-insensitive):
  date, description, amount, category, type

OFX:
  Basic support for <STMTTRN> blocks with DTPOSTED, TRNAMT, MEMO.`,
		RunE: runImport,
	}

	importCmd.Flags().StringVarP(&importFile, "file", "f", "", "Path to CSV or OFX file")
	_ = importCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(importCmd)
}

func runImport(cmd *cobra.Command, args []string) error {
	ext := strings.ToLower(filepath.Ext(importFile))

	database, err := db.Open("finance.db")
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}
	defer database.Close()

	trRepo := repo.NewTransactionRepo(database)
	ruleRepo := repo.NewRuleRepo(database)

	switch ext {
	case ".csv":
		return importCSV(importFile, trRepo, ruleRepo)
	case ".ofx":
		return importOFX(importFile, trRepo, ruleRepo)
	default:
		return fmt.Errorf("unsupported file extension: %s (use .csv or .ofx)", ext)
	}
}

func importCSV(path string, trRepo *repo.TransactionRepo, ruleRepo *repo.RuleRepo) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.TrimLeadingSpace = true
	// for regional settings where ; is used as separator
	reader.Comma = ';'

	// Read header
	header, err := reader.Read()
	if err != nil {
		return fmt.Errorf("read header: %w", err)
	}

	// If the whole header is in one cell like "date,description,...", split manually.
	if len(header) == 1 && strings.Contains(header[0], ",") {
		raw := strings.TrimSpace(header[0])
		raw = strings.Trim(raw, "\"")
		header = strings.Split(raw, ",")
	}

	// Map column name -> index (lowercased)
	colIndex := map[string]int{}
	for i, h := range header {
		name := strings.ToLower(strings.TrimSpace(h))
		colIndex[name] = i
	}

	// Required: description, amount
	descIdx, ok := colIndex["description"]
	if !ok {
		return fmt.Errorf("CSV must contain a 'description' column")
	}
	amountIdx, ok := colIndex["amount"]
	if !ok {
		return fmt.Errorf("CSV must contain an 'amount' column")
	}

	// Optional
	dateIdx, hasDate := colIndex["date"]
	categoryIdx, hasCategory := colIndex["category"]
	typeIdx, hasType := colIndex["type"]

	var imported, failed int

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("read row: %w", err)
		}

		description := strings.TrimSpace(record[descIdx])
		if description == "" {
			failed++
			continue
		}

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

		// If category empty, try rules
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

		if err := trRepo.Insert(tx); err != nil {
			fmt.Printf("Failed to insert transaction '%s': %v\n", description, err)
			failed++
			continue
		}

		imported++
	}

	fmt.Printf("CSV import finished. Imported: %d, Skipped: %d\n", imported, failed)
	return nil
}

func importOFX(path string, trRepo *repo.TransactionRepo, ruleRepo *repo.RuleRepo) error {
	txs, err := service.ParseOFX(path)
	if err != nil {
		return fmt.Errorf("parse OFX: %w", err)
	}

	if len(txs) == 0 {
		fmt.Println("No transactions found in OFX file.")
		return nil
	}

	var imported, failed int

	for i := range txs {
		tx := &txs[i]

		// try auto category via rules if empty
		if tx.Category == "" {
			if cat, err := ruleRepo.FindCategory(tx.Description); err == nil && cat != "" {
				tx.Category = cat
			}
		}

		if err := trRepo.Insert(tx); err != nil {
			fmt.Printf("Failed to insert OFX transaction '%s': %v\n", tx.Description, err)
			failed++
			continue
		}

		imported++
	}

	fmt.Printf("OFX import finished. Imported: %d, Skipped: %d\n", imported, failed)
	return nil
}
