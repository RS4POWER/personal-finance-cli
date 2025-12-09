package repo

import (
	"database/sql"

	"github.com/RS4POWER/personal-finance-cli/internal/domain"
)

type BudgetRepo struct {
	db *sql.DB
}

func NewBudgetRepo(db *sql.DB) *BudgetRepo {
	return &BudgetRepo{db: db}
}

func (r *BudgetRepo) Insert(b *domain.Budget) error {
	res, err := r.db.Exec(`
        INSERT INTO budgets (category, limit_amount)
        VALUES (?, ?)`,
		b.Category,
		b.Limit,
	)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	b.ID = id
	return nil
}

func (r *BudgetRepo) List() ([]domain.Budget, error) {
	rows, err := r.db.Query(`
        SELECT id, category, limit_amount
        FROM budgets
        ORDER BY category`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Budget
	for rows.Next() {
		var b domain.Budget
		if err := rows.Scan(&b.ID, &b.Category, &b.Limit); err != nil {
			return nil, err
		}
		result = append(result, b)
	}

	return result, rows.Err()
}
