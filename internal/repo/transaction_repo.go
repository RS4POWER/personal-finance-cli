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
