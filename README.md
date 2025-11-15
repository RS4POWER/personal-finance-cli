# Project: Personal Finance CLI Manager

## Description

This project is a command-line tool for tracking personal income and expenses.  
The goal is to be able to import transactions from bank statements, categorize them automatically, set budgets, and generate insightful reports — all from the terminal.

At this checkpoint, the architecture is finalized and basic functionality is implemented (add/search/report using a SQLite database).

---

## User Stories

| Feature / User Story                                            | Status | Notes |
|-----------------------------------------------------------------|--------|-------|
| As a user, I can import transactions from CSV/OFX files         | ❌     | planned for next iterations |
| As a user, I can manually add income and expenses               | ✅     | `pfcli add` inserts into SQLite |
| As a user, I can categorize transactions automatically          | ❌     | will be based on regex rules |
| As a user, I can set budgets per category and get alerts        | ❌     | CLI stub exists (`pfcli budget`) |
| As a user, I can generate reports (monthly spending, breakdown) | ⚠️     | basic summary via `pfcli report` |
| As a user, I can search and filter transactions                 | ✅     | simple text search via `pfcli search` |

---

## Usage

The main CLI entrypoint is `pfcli`.

### Available commands

| Command                                   | Description |
|-------------------------------------------|-------------|
| `pfcli add`                               | Add a transaction manually (income or expense) |
| `pfcli search --text <query>`             | Search transactions by text in the description |
| `pfcli report`                            | Show a simple summary: total income, expense, balance |
| `pfcli import`                            | Stub: will import transactions from CSV/OFX files |
| `pfcli budget`                            | Stub: will manage budgets and alerts per category |
| `pfcli tui`                               | Stub: will start an interactive terminal UI |

### Examples

➕ Add a transaction

pfcli add --amount 25.5 --description "Pizza" --category "Food" --type expense
pfcli add --amount 5000 --description "Salary November" --category "Income" --type income

🔍 Search transactions

pfcli search --text Pizza

📊 Generate financial report

pfcli report


## System Architecture

The project uses a layered architecture with clear separation between CLI, services, repositories, database access and domain models.

 ```mermaid

graph TD
    User[User / Terminal] -->|CLI commands: add, search, report, import, budget, tui| CLI[Cobra CLI (cmd/pfcli + internal/cli)]

    CLI --> Services[Service Layer (internal/service)]
    Services --> Repo[Repository Layer (internal/repo)]
    Repo --> DB[Database Layer - SQLite (internal/db)]
    DB --> Repo

    Services --> Domain[Domain Models (internal/domain)]
    Repo --> Domain

    %% Future service responsibilities
    Services --> ImportService[Import Service (CSV/OFX)]
    Services --> Categorization[Categorization (regex rules)]
    Services --> BudgetManager[Budget Manager & Alerts]
    Services --> ReportGen[Advanced Reports & Charts]
    Services --> TUIBackend[TUI Backend Logic]
```
## Project Structure

personal-finance-cli/
├── cmd/
│   └── pfcli/
│       └── main.go            # CLI entrypoint
│
├── internal/
│   ├── cli/                   # Cobra commands (add, search, report, import, budget, tui)
│   ├── db/                    # SQLite connection + migrations
│   ├── domain/                # Domain entities (Transaction, etc.)
│   └── repo/                  # TransactionRepo (Insert, SearchByText, Totals)
│
└── docs/
    └── architecture.md        # Detailed architecture description


## 🏛️ Architecture Overview

### 🟦 1. CLI Layer (internal/cli)

Handles:

parsing user input

subcommands (add, search, report, etc.)

validation
Framework: Cobra

### 🟩 2. Service Layer (future work)

Will handle:

categorization rules

budgets & alerts

business logic
(Currently empty by design)

### 🟧 3. Repository Layer (internal/repo)

Implements:

Insert()

SearchByText()

Totals()

Keeps SQL logic isolated from CLI.

### 🟪 4. Database Layer (internal/db)

Responsibilities:

SQLite database initialization

automatic migrations

persistent finance.db file

Driver: glebarez/sqlite (CGO-free)

### 🟨 5. Domain Layer (internal/domain)

Contains pure business objects:

Transaction

TransactionTypeIncome

TransactionTypeExpense


## Notes (Checkpoint)

System architecture is designed and fixed (folders, layers, responsibilities).

Basic functionality is implemented and working:

add (write to SQLite)

search (read & filter)

report (aggregate totals)

Stub commands (import, budget, tui) exist to match project requirements and will be implemented in next iterations.


## 🚀 Running the Project

### Clone repository:

git clone https://github.com/RS4POWER/personal-finance-cli
cd personal-finance-cli

### Run CLI:
go run ./cmd/pfcli