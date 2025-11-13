#### \# Architecture



This project is a personal finance CLI manager written in Go. It tracks income and expenses using a local SQLite database and a command-line interface built with Cobra.



The goal of the application is to support the following features:



\- Import transactions from CSV/OFX files.

\- Manually add income and expenses.

\- Automatically categorize transactions using rules (regex).

\- Set budgets per category and get alerts.

\- Generate reports (monthly spending, category breakdown, charts).

\- Search and filter transactions.

\- Browse transactions via an interactive TUI.



!!!! At this checkpoint, only part of the functionality is implemented, but the architecture is designed to support all of the above. !!!!



---



#### \## High-level structure



Project layout:



```text

cmd/pfcli

&nbsp; main.go           # Entry point of the CLI application



internal/

&nbsp; cli/              # CLI layer (Cobra commands)

&nbsp;   root.go

&nbsp;   add.go

&nbsp;   search.go

&nbsp;   report.go

&nbsp;   import.go       # stub

&nbsp;   budget.go       # stub

&nbsp;   tui.go          # stub



&nbsp; db/               # Database connection and migrations

&nbsp;   db.go



&nbsp; domain/           # Domain models

&nbsp;   transaction.go

&nbsp;   budget.go

&nbsp;   category\_rule.go



&nbsp; repo/             # Data access layer (repositories)

&nbsp;   transaction\_repo.go

&nbsp;   budget\_repo.go      # stub

&nbsp;   rule\_repo.go        # stub



&nbsp; service/          # Business logic layer (planned to grow)

&nbsp;   import\_service.go   # stub

&nbsp;   categorization.go   # stub

&nbsp;   budgets.go          # stub

&nbsp;   reports.go          # stub

&nbsp;   tui\_service.go      # stub



---



#### Layers



CLI (internal/cli)



Uses github.com/spf13/cobra for subcommands: add, search, report (implemented), and import, budget, tui (currently stubs).



Responsible only for:



Parsing command-line arguments and flags.



Printing results to the terminal.



Calling the appropriate services/repositories.



DB (internal/db)



Provides a single function Open(path string) (\*sql.DB, error).



Uses the github.com/glebarez/sqlite driver to open a SQLite database.



Runs migrations on startup to ensure that required tables are created (currently the transactions table).



Domain (internal/domain)



Contains the business entities (pure Go structs):



Transaction – represents a single financial transaction.



TransactionType – income or expense.



Budget – planned amount per category.



CategoryRule – regex-based rules for automatic categorization.



This layer is independent of the database and CLI.



Repositories (internal/repo)



Wrap \*sql.DB and provide high-level operations:



TransactionRepo:



Insert – add new transactions.



SearchByText – filter by description.



Totals – compute income, expense and balance.



BudgetRepo (planned).



RuleRepo (planned).



Responsible for translating between SQL rows and domain models.



Services (internal/service)



Orchestrate higher-level use cases:



import\_service.go – import transactions from CSV/OFX files.



categorization.go – apply regex rules to assign categories automatically.



budgets.go – check budgets and generate alerts.



reports.go – generate monthly/category reports and charts for the terminal.



tui\_service.go – backend logic for the interactive terminal UI.



####   **!! At this checkpoint, service files exist as stubs and will be implemented later. !!**





**---**



#### **Data flow examples**



**Add transaction**



**User runs:**



**pfcli add --amount 25.5 --description "Pizza" --category Food --type expense**





**internal/cli/add.go parses CLI flags and builds a domain.Transaction.**



**It opens the database via internal/db.Open("finance.db").**



**It creates a TransactionRepo with the \*sql.DB.**



**It calls TransactionRepo.Insert to insert the transaction into the transactions table.**



**It prints the assigned transaction ID to the terminal.**





