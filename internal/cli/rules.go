package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/RS4POWER/personal-finance-cli/internal/db"
	"github.com/RS4POWER/personal-finance-cli/internal/domain"
	"github.com/RS4POWER/personal-finance-cli/internal/repo"
)

var (
	rulePattern  string
	ruleCategory string
)

func init() {
	rulesCmd := &cobra.Command{
		Use:   "rules",
		Short: "Manage automatic categorization rules (regex → category)",
	}

	addRuleCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new categorization rule",
		RunE:  runAddRule,
	}
	addRuleCmd.Flags().StringVar(&rulePattern, "pattern", "", "Regex pattern to match in description")
	addRuleCmd.Flags().StringVar(&ruleCategory, "category", "", "Category to assign when the pattern matches")
	_ = addRuleCmd.MarkFlagRequired("pattern")
	_ = addRuleCmd.MarkFlagRequired("category")

	listRulesCmd := &cobra.Command{
		Use:   "list",
		Short: "List all categorization rules",
		RunE:  runListRules,
	}

	rulesCmd.AddCommand(addRuleCmd)
	rulesCmd.AddCommand(listRulesCmd)

	rootCmd.AddCommand(rulesCmd)
}

func runAddRule(cmd *cobra.Command, args []string) error {
	database, err := db.Open("finance.db")
	if err != nil {
		return err
	}
	defer database.Close()

	r := repo.NewRuleRepo(database)

	rule := &domain.CategoryRule{
		Pattern:  rulePattern,
		Category: ruleCategory,
	}

	if err := r.Insert(rule); err != nil {
		return err
	}

	fmt.Printf("Rule added with ID %d: /%s/ -> %s\n", rule.ID, rule.Pattern, rule.Category)
	return nil
}

func runListRules(cmd *cobra.Command, args []string) error {
	database, err := db.Open("finance.db")
	if err != nil {
		return err
	}
	defer database.Close()

	r := repo.NewRuleRepo(database)

	rules, err := r.List()
	if err != nil {
		return err
	}

	if len(rules) == 0 {
		fmt.Println("No rules defined.")
		return nil
	}

	for _, rule := range rules {
		fmt.Printf("[%d] /%s/ -> %s\n", rule.ID, rule.Pattern, rule.Category)
	}

	return nil
}
