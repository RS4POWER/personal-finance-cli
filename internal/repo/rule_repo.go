package repo

import (
	"database/sql"
	"regexp"

	"github.com/RS4POWER/personal-finance-cli/internal/domain"
)

type RuleRepo struct {
	db *sql.DB
}

func NewRuleRepo(db *sql.DB) *RuleRepo {
	return &RuleRepo{db: db}
}

func (r *RuleRepo) Insert(rule *domain.CategoryRule) error {
	res, err := r.db.Exec(`
        INSERT INTO category_rules (pattern, category)
        VALUES (?, ?)`,
		rule.Pattern,
		rule.Category,
	)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	rule.ID = id
	return nil
}

func (r *RuleRepo) List() ([]domain.CategoryRule, error) {
	rows, err := r.db.Query(`
        SELECT id, pattern, category
        FROM category_rules
        ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.CategoryRule
	for rows.Next() {
		var rule domain.CategoryRule
		if err := rows.Scan(&rule.ID, &rule.Pattern, &rule.Category); err != nil {
			return nil, err
		}
		result = append(result, rule)
	}

	return result, rows.Err()
}

// FindCategory returns the first matching category for the given description,
// or empty string if no regex rule matches.
func (r *RuleRepo) FindCategory(description string) (string, error) {
	rules, err := r.List()
	if err != nil {
		return "", err
	}

	for _, rule := range rules {
		re, err := regexp.Compile("(?i)" + rule.Pattern) // (?i) = case-insensitive
		if err != nil {
			// ignore invalid regex rules
			continue
		}
		if re.MatchString(description) {
			return rule.Category, nil
		}
	}

	return "", nil
}
