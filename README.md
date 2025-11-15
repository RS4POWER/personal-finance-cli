# <div align="center">

# 

# \# 🚀 \*\*Personal Finance CLI Manager\*\*

# 

# A modern, modular and extensible \*\*command-line application\*\* for tracking income and expenses.<div align="center">

# 

# \# 🚀 \*\*Personal Finance CLI Manager\*\*

# 

# A modern, modular and extensible \*\*command-line application\*\* for tracking income and expenses.

# Built in \*\*Go\*\*, using \*\*SQLite\*\*, \*\*Cobra\*\*, and clean layered architecture.

# 

# \### 🎓 \*University Project\*

# 

# \*\*Advanced Technologies for Application Development\*\*

# UPT – CTI Master, Year 2

# 

# ---

# 

# !\[Go](https://img.shields.io/badge/Go-1.25-blue?logo=go)

# !\[Platform](https://img.shields.io/badge/Platform-Windows%2011-lightgrey?logo=windows)

# !\[Status](https://img.shields.io/badge/Status-Checkpoint%20Complete-brightgreen)

# !\[Architecture](https://img.shields.io/badge/Architecture-Layered-blueviolet)

# !\[SQLite](https://img.shields.io/badge/DB-SQLite-orange?logo=sqlite)

# 

# </div>

# 

# ---

# 

# \# 📌 \*\*Features (Checkpoint Status)\*\*

# 

# | User Story                       | Status | Notes                                |

# | -------------------------------- | ------ | ------------------------------------ |

# | Import CSV/OFX                   | ❌      | Stub available                       |

# | Manual add of transactions       | ✅      | Fully implemented                    |

# | Automatic categorization (regex) | ❌      | Planned                              |

# | Budgets + alerts                 | ❌      | Stub available                       |

# | Monthly / Category reports       | ⚠️     | Basic version implemented (`report`) |

# | Search \& filtering               | ✅      | Text search functional               |

# | Interactive TUI                  | ❌      | Stub included                        |

# 

# ---

# 

# \# 📂 \*\*Project Structure\*\*

# 

# ```

# personal-finance-cli/

# │

# ├── cmd/

# │   └── pfcli/

# │       └── main.go                 # Entry point

# │

# ├── internal/

# │   ├── cli/                        # CLI commands (Cobra)

# │   │   ├── root.go

# │   │   ├── add.go

# │   │   ├── search.go

# │   │   ├── report.go

# │   │   ├── import.go

# │   │   ├── budget.go

# │   │   └── tui.go

# │   │

# │   ├── db/

# │   │   └── db.go                   # SQLite connection + migrations

# │   │

# │   ├── domain/

# │   │   └── transaction.go          # Domain entity

# │   │

# │   └── repo/

# │       └── transaction\_repo.go     # DB access (Insert, Search, Totals)

# │

# ├── docs/

# │   └── architecture.md             # Full architecture description

# │

# ├── go.mod

# ├── go.sum

# └── README.md

# ```

# 

# ---

# 

# \# 🏛️ \*\*Architecture Overview\*\*

# 

# \### 🟦 \*\*1. CLI Layer (`internal/cli`)\*\*

# 

# Handles:

# 

# \* parsing user input

# \* subcommands (`add`, `search`, `report`, etc.)

# \* validation

# &nbsp; Framework: \*\*Cobra\*\*

# 

# ---

# 

# \### 🟩 \*\*2. Service Layer (future work)\*\*

# 

# Will handle:

# 

# \* categorization rules

# \* budgets \& alerts

# \* business logic

# &nbsp; (Currently empty by design)

# 

# ---

# 

# \### 🟧 \*\*3. Repository Layer (`internal/repo`)\*\*

# 

# Implements:

# 

# \* `Insert()`

# \* `SearchByText()`

# \* `Totals()`

# 

# Keeps SQL logic isolated from CLI.

# 

# ---

# 

# \### 🟪 \*\*4. Database Layer (`internal/db`)\*\*

# 

# Responsibilities:

# 

# \* SQLite database initialization

# \* automatic migrations

# \* persistent `finance.db` file

# 

# Driver: \*\*glebarez/sqlite\*\* (CGO-free)

# 

# ---

# 

# \### 🟨 \*\*5. Domain Layer (`internal/domain`)\*\*

# 

# Contains pure business objects:

# 

# \* `Transaction`

# \* `TransactionTypeIncome`

# \* `TransactionTypeExpense`

# 

# ---

# 

# \# 🖥️ \*\*Usage\*\*

# 

# \### ➕ Add a transaction

# 

# ```bash

# pfcli add --amount 25.5 --description "Pizza" --category "Food" --type expense

# ```

# 

# ---

# 

# \### 🔍 Search transactions

# 

# ```bash

# pfcli search --text Pizza

# ```

# 

# Output example:

# 

# ```

# \[2] 2025-11-13 | expense | 60.00 | Food | Pizza with friends

# ```

# 

# ---

# 

# \### 📊 Generate financial report

# 

# ```bash

# pfcli report

# ```

# 

# Example:

# 

# ```

# Income : 5200.00

# Expense: 210.00

# Balance: 4990.00

# ```

# 

# ---

# 

# \### 📁 Import CSV/OFX \*(stub)\*

# 

# ```bash

# pfcli import

# ```

# 

# \### 💰 Budgets \*(stub)\*

# 

# ```bash

# pfcli budget

# ```

# 

# \### 🖼️ TUI Interface \*(stub)\*

# 

# ```bash

# pfcli tui

# ```

# 

# ---

# 

# \# 🔧 \*\*Tech Stack\*\*

# 

# | Component     | Technology               |

# | ------------- | ------------------------ |

# | Language      | Go 1.25                  |

# | Database      | SQLite (glebarez/sqlite) |

# | CLI Framework | spf13/cobra              |

# | Architecture  | Layered / Modular        |

# | OS            | Windows 11 (development) |

# 

# ---

# 

# \# 🚀 \*\*Running the Project\*\*

# 

# Clone repository:

# 

# ```bash

# git clone https://github.com/RS4POWER/personal-finance-cli

# cd personal-finance-cli

# ```

# 

# Run CLI:

# 

# ```bash

# go run ./cmd/pfcli

# ```

# 

# ---

# 

# \# 📝 \*\*Future Enhancements\*\*

# 

# \* \[ ] Import CSV/OFX

# \* \[ ] Regex-based classification

# \* \[ ] Budget limits + alerts

# \* \[ ] Monthly ASCII charts

# \* \[ ] Full TUI mode

# \* \[ ] Export JSON/CSV

# \* \[ ] Unit tests

# 

# ---

# 

# \# 🎯 \*\*Checkpoint Summary\*\*

# 

# | Requirement            | Status                      |

# | ---------------------- | --------------------------- |

# | System Architecture    | ✔ Completed                 |

# | Basic Functionality    | ✔ `add`, `search`, `report` |

# | CLI with subcommands   | ✔ Fully set up              |

# | Partial Implementation | ✔ Delivered                 |

# | Documentation          | ✔ README + architecture.md  |

# 

# ---

# 

# <div align="center">

# 

# \### 💙 \*Project ready for Checkpoint Submission\*

# 

# \*\*Made with Go, coffee, and debugging power.\*\*

# 

# </div>



# Built in \*\*Go\*\*, using \*\*SQLite\*\*, \*\*Cobra\*\*, and clean layered architecture.

# 

# \### 🎓 \*University Project\*

# 

# \*\*Advanced Technologies for Application Development\*\*

# UPT – CTI Master, Year 2

# 

# ---

# 

# !\[Go](https://img.shields.io/badge/Go-1.25-blue?logo=go)

# !\[Platform](https://img.shields.io/badge/Platform-Windows%2011-lightgrey?logo=windows)

# !\[Status](https://img.shields.io/badge/Status-Checkpoint%20Complete-brightgreen)

# !\[Architecture](https://img.shields.io/badge/Architecture-Layered-blueviolet)

# !\[SQLite](https://img.shields.io/badge/DB-SQLite-orange?logo=sqlite)

# 

# </div>

# 

# ---

# 

# \# 📌 \*\*Features (Checkpoint Status)\*\*

# 

# | User Story                       | Status | Notes                                |

# | -------------------------------- | ------ | ------------------------------------ |

# | Import CSV/OFX                   | ❌      | Stub available                       |

# | Manual add of transactions       | ✅      | Fully implemented                    |

# | Automatic categorization (regex) | ❌      | Planned                              |

# | Budgets + alerts                 | ❌      | Stub available                       |

# | Monthly / Category reports       | ⚠️     | Basic version implemented (`report`) |

# | Search \& filtering               | ✅      | Text search functional               |

# | Interactive TUI                  | ❌      | Stub included                        |

# 

# ---

# 

# \# 📂 \*\*Project Structure\*\*

# 

# ```

# personal-finance-cli/

# │

# ├── cmd/

# │   └── pfcli/

# │       └── main.go                 # Entry point

# │

# ├── internal/

# │   ├── cli/                        # CLI commands (Cobra)

# │   │   ├── root.go

# │   │   ├── add.go

# │   │   ├── search.go

# │   │   ├── report.go

# │   │   ├── import.go

# │   │   ├── budget.go

# │   │   └── tui.go

# │   │

# │   ├── db/

# │   │   └── db.go                   # SQLite connection + migrations

# │   │

# │   ├── domain/

# │   │   └── transaction.go          # Domain entity

# │   │

# │   └── repo/

# │       └── transaction\_repo.go     # DB access (Insert, Search, Totals)

# │

# ├── docs/

# │   └── architecture.md             # Full architecture description

# │

# ├── go.mod

# ├── go.sum

# └── README.md

# ```

# 

# ---

# 

# \# 🏛️ \*\*Architecture Overview\*\*

# 

# \### 🟦 \*\*1. CLI Layer (`internal/cli`)\*\*

# 

# Handles:

# 

# \* parsing user input

# \* subcommands (`add`, `search`, `report`, etc.)

# \* validation

# &nbsp; Framework: \*\*Cobra\*\*

# 

# ---

# 

# \### 🟩 \*\*2. Service Layer (future work)\*\*

# 

# Will handle:

# 

# \* categorization rules

# \* budgets \& alerts

# \* business logic

# &nbsp; (Currently empty by design)

# 

# ---

# 

# \### 🟧 \*\*3. Repository Layer (`internal/repo`)\*\*

# 

# Implements:

# 

# \* `Insert()`

# \* `SearchByText()`

# \* `Totals()`

# 

# Keeps SQL logic isolated from CLI.

# 

# ---

# 

# \### 🟪 \*\*4. Database Layer (`internal/db`)\*\*

# 

# Responsibilities:

# 

# \* SQLite database initialization

# \* automatic migrations

# \* persistent `finance.db` file

# 

# Driver: \*\*glebarez/sqlite\*\* (CGO-free)

# 

# ---

# 

# \### 🟨 \*\*5. Domain Layer (`internal/domain`)\*\*

# 

# Contains pure business objects:

# 

# \* `Transaction`

# \* `TransactionTypeIncome`

# \* `TransactionTypeExpense`

# 

# ---

# 

# \# 🖥️ \*\*Usage\*\*

# 

# \### ➕ Add a transaction

# 

# ```bash

# pfcli add --amount 25.5 --description "Pizza" --category "Food" --type expense

# ```

# 

# ---

# 

# \### 🔍 Search transactions

# 

# ```bash

# pfcli search --text Pizza

# ```

# 

# Output example:

# 

# ```

# \[2] 2025-11-13 | expense | 60.00 | Food | Pizza with friends

# ```

# 

# ---

# 

# \### 📊 Generate financial report

# 

# ```bash

# pfcli report

# ```

# 

# Example:

# 

# ```

# Income : 5200.00

# Expense: 210.00

# Balance: 4990.00

# ```

# 

# ---

# 

# \### 📁 Import CSV/OFX \*(stub)\*

# 

# ```bash

# pfcli import

# ```

# 

# \### 💰 Budgets \*(stub)\*

# 

# ```bash

# pfcli budget

# ```

# 

# \### 🖼️ TUI Interface \*(stub)\*

# 

# ```bash

# pfcli tui

# ```

# 

# ---

# 

# \# 🔧 \*\*Tech Stack\*\*

# 

# | Component     | Technology               |

# | ------------- | ------------------------ |

# | Language      | Go 1.25                  |

# | Database      | SQLite (glebarez/sqlite) |

# | CLI Framework | spf13/cobra              |

# | Architecture  | Layered / Modular        |

# | OS            | Windows 11 (development) |

# 

# ---

# 

# \# 🚀 \*\*Running the Project\*\*

# 

# Clone repository:

# 

# ```bash

# git clone https://github.com/RS4POWER/personal-finance-cli

# cd personal-finance-cli

# ```

# 

# Run CLI:

# 

# ```bash

# go run ./cmd/pfcli

# ```

# 

# ---

# 

# \# 📝 \*\*Future Enhancements\*\*

# 

# \* \[ ] Import CSV/OFX

# \* \[ ] Regex-based classification

# \* \[ ] Budget limits + alerts

# \* \[ ] Monthly ASCII charts

# \* \[ ] Full TUI mode

# \* \[ ] Export JSON/CSV

# \* \[ ] Unit tests

# 

# ---

# 

# \# 🎯 \*\*Checkpoint Summary\*\*

# 

# | Requirement            | Status                      |

# | ---------------------- | --------------------------- |

# | System Architecture    | ✔ Completed                 |

# | Basic Functionality    | ✔ `add`, `search`, `report` |

# | CLI with subcommands   | ✔ Fully set up              |

# | Partial Implementation | ✔ Delivered                 |

# | Documentation          | ✔ README + architecture.md  |

# 

# ---

# 

# <div align="center">

# 

# \### 💙 \*Project ready for Checkpoint Submission\*

# 

# \*\*Made with Go, coffee, and debugging power.\*\*

# 

# </div>



