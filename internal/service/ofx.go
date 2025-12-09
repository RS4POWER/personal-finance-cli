package service

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/RS4POWER/personal-finance-cli/internal/domain"
)

func ParseOFX(path string) ([]domain.Transaction, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	text := string(data)

	// Very simple OFX parsing:
	// We look for <STMTTRN> blocks and extract <DTPOSTED>, <TRNAMT>, <MEMO>
	blockRe := regexp.MustCompile(`(?s)<STMTTRN>(.*?)</STMTTRN>`)
	dateRe := regexp.MustCompile(`<DTPOSTED>(\d{8})`)
	amountRe := regexp.MustCompile(`<TRNAMT>([-0-9.]+)`)
	memoRe := regexp.MustCompile(`<MEMO>([^<\r\n]+)`)

	var result []domain.Transaction

	blocks := blockRe.FindAllStringSubmatch(text, -1)
	for _, b := range blocks {
		block := b[1]

		dateStr := ""
		if m := dateRe.FindStringSubmatch(block); len(m) == 2 {
			dateStr = m[1]
		}
		amountStr := ""
		if m := amountRe.FindStringSubmatch(block); len(m) == 2 {
			amountStr = m[1]
		}
		memo := ""
		if m := memoRe.FindStringSubmatch(block); len(m) == 2 {
			memo = strings.TrimSpace(m[1])
		}

		if amountStr == "" || memo == "" {
			continue
		}

		amt, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			continue
		}

		// parse date YYYYMMDD
		var date time.Time
		if dateStr != "" {
			d, err := time.Parse("20060102", dateStr)
			if err == nil {
				date = d
			} else {
				date = time.Now()
			}
		} else {
			date = time.Now()
		}

		// Heuristic: positive = income, negative = expense
		tType := domain.TransactionTypeExpense
		if amt > 0 {
			tType = domain.TransactionTypeIncome
		} else {
			amt = -amt
			tType = domain.TransactionTypeExpense
		}

		result = append(result, domain.Transaction{
			Date:        date,
			Description: memo,
			Amount:      amt,
			Category:    "", // can be filled by rules later
			Type:        tType,
		})
	}

	return result, nil
}
