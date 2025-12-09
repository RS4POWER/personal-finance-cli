package repo

import (
	"database/sql"
	"time"

	"github.com/RS4POWER/personal-finance-cli/internal/domain"
)

type TransactionRepo struct {
	db *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{db: db}
}

func (r *TransactionRepo) Insert(t *domain.Transaction) error {
	res, err := r.db.Exec(`
        INSERT INTO transactions (date, description, amount, category, type)
        VALUES (?, ?, ?, ?, ?)`,
		t.Date.Format("2006-01-02"),
		t.Description,
		t.Amount,
		t.Category,
		string(t.Type),
	)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = id
	return nil
}

func (r *TransactionRepo) SearchByText(text string) ([]domain.Transaction, error) {
	rows, err := r.db.Query(`
        SELECT id, date, description, amount, category, type
        FROM transactions
        WHERE description LIKE ?
        ORDER BY date`, "%"+text+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Transaction
	for rows.Next() {
		var t domain.Transaction
		var dateStr, typeStr string
		if err := rows.Scan(&t.ID, &dateStr, &t.Description, &t.Amount, &t.Category, &typeStr); err != nil {
			return nil, err
		}
		t.Date, _ = time.Parse("2006-01-02", dateStr)
		t.Type = domain.TransactionType(typeStr)
		result = append(result, t)
	}
	return result, rows.Err()
}

func (r *TransactionRepo) Totals() (income, expense float64, err error) {
	rows, err := r.db.Query(`
        SELECT type, SUM(amount)
        FROM transactions
        GROUP BY type`)
	if err != nil {
		return 0, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var t string
		var sum float64

		if err := rows.Scan(&t, &sum); err != nil {
			return 0, 0, err
		}

		if t == string(domain.TransactionTypeIncome) {
			income = sum
		} else if t == string(domain.TransactionTypeExpense) {
			expense = sum
		}
	}

	return income, expense, rows.Err()
}

func (r *TransactionRepo) TotalsByCategory(t domain.TransactionType) ([]domain.CategoryTotal, error) {
	rows, err := r.db.Query(`
        SELECT COALESCE(category, ''), SUM(amount)
        FROM transactions
        WHERE type = ?
        GROUP BY category
        ORDER BY SUM(amount) DESC
    `, string(t))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.CategoryTotal
	for rows.Next() {
		var c domain.CategoryTotal
		if err := rows.Scan(&c.Category, &c.Total); err != nil {
			return nil, err
		}
		result = append(result, c)
	}
	return result, rows.Err()
}

func (r *TransactionRepo) LastN(n int) ([]domain.Transaction, error) {
	rows, err := r.db.Query(`
        SELECT id, date, description, amount, category, type
        FROM transactions
        ORDER BY date DESC, id DESC
        LIMIT ?`, n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Transaction
	for rows.Next() {
		var t domain.Transaction
		var dateStr, typeStr string
		if err := rows.Scan(&t.ID, &dateStr, &t.Description, &t.Amount, &t.Category, &typeStr); err != nil {
			return nil, err
		}
		t.Date, _ = time.Parse("2006-01-02", dateStr)
		t.Type = domain.TransactionType(typeStr)
		result = append(result, t)
	}
	return result, rows.Err()
}
