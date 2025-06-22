# Fiber Boilerplate with Air Hot Reloading

A production-ready GoFiber web application boilerplate following best practices from the .clinerules specification.

## Project Structure

```mermaid
graph TD
    A[Project Root] --> B[main.go]
    A --> C[routes/]
    A --> D[handlers/]
    A --> E[views/]
    A --> F[middlewares/]
    A --> G[configs/]
    A --> H[static/]
    A --> I[Makefile]
    E --> J[layouts/]
    E --> K[pages/]
    E --> L[partials/]
```

## Features

- **Core Framework**: [Fiber](https://github.com/gofiber/fiber) - High performance HTTP framework
- **Development Tools**:
  - [Air](https://github.com/cosmtrek/air) - Hot reloading (.air.toml config)
  - Built-in Makefile workflow
  - dev.sh helper scripts
- **Authentication**:
  - Login/Register handlers
  - Session management
- **Admin Dashboard**:
  - Dedicated admin layout
  - Protected routes
- **Architecture**:
  - Route grouping and separation
  - Handler/store pattern for database abstraction
  - Custom middleware support
  - Configuration management (configs/)
- **Template Engine**:
  - Go's html/template integrated via Fiber adapter
  - Layouts (main.html, admin.html) and partials support
- **Static Assets**:
  - CSS framework integration (Bootstrap/Tailwind)
  - Organized in static/ directory
- **Operational**:
  - Health check endpoints (/livez, /readyz)
  - Metrics dashboard (/metrics)
  - Graceful shutdown

## Requirements

- Go 1.18+
- Air (for development)
- golangci-lint (for code quality)

## Getting Started

### Installation

```bash
git clone https://github.com/yourusername/fiber-boilerplate.git
cd fiber-boilerplate
go mod download
```

### Development Workflow

```mermaid
graph LR
    A[Code Changes] --> B[Air Reload]
    B --> C[Automatic Build]
    C --> D[Restart Server]
```

Key commands:
```bash
make dev    # Run with hot reloading (uses Air)
make build  # Production build
make test   # Run tests
make lint   # Run linters
```

## Configuration

Configure via `.env` file:

| Variable          | Default            | Description                     |
|-------------------|--------------------|---------------------------------|
| PORT              | 3000               | Server port                    |
| APP_NAME          | "Fiber Boilerplate"| Application name               |
| SHUTDOWN_TIMEOUT  | 5                  | Graceful shutdown timeout (sec)|

## File Structure Details

```markdown
/
├── .air.toml           # Air configuration
├── .clinerules         # Custom assistant instructions
├── .env.example        # Environment template
├── .gitignore
├── dev.sh              # Development utilities
├── go.mod
├── go.sum
├── main.go             # Main entry point
├── Makefile
├── README.md
├── template.html       # HTML template reference
├── configs/            # Configuration management
│   └── configs.go
├── handlers/           # HTTP handlers
│   ├── auth_handler.go
│   ├── dashboard_handler.go
│   └── home_handler.go
├── middlewares/        # Custom middleware
├── models/             # Data structures (placeholder)
├── routes/             # Route definitions
│   └── routes.go
├── static/             # Static assets
│   └── css/
│       └── styles.css
├── store/              # Database layer (placeholder)
└── views/              # Templates
    ├── layouts/
    │   ├── admin.html
    │   └── main.html
    ├── pages/
    │   ├── dashboard.html
    │   ├── home.html
    │   ├── login.html
    │   └── register.html
    └── partials/       # Reusable components
```

## Architecture Overview

```mermaid
graph LR
    R[Routes] --> M[Middlewares]
    M --> H[Handlers]
    H --> S[Store]
    H --> C[Configs]
    S --> DB[(Database)]
    H --> V[Views]
    V --> A[Assets]
```

### Key Patterns

1. **Route Definitions**:
   - Grouped by resource/functionality
   - Defined in `routes/` package
   - Clean separation from handlers

2. **Handler/Store Pattern**:
   - Handlers process HTTP requests
   - Store handles database operations
   - Clear separation of concerns

3. **Configuration Management**:
   - Centralized in configs/ package
   - Environment variable support

4. **Authentication Flow**:
   - Login/Register handlers
   - Session management
   - Protected routes

5. **Template Engine**:
   - Configured in `main.go`
   - Supports multiple layouts (`views/layouts/`)
   - Pages in `views/pages/`
   - Reusable partials (`views/partials/`)

## Code Standards

- Follows "Effective Go" idioms
- Strict error handling (never ignore errors)
- Linting via golangci-lint
- Formatted with go fmt
- Validation for all incoming data
- Configuration via .clinerules

## Air Configuration

Customize hot-reloading in `.air.toml`:

```toml
[build]
  delay = 1000               # Delay after changes (ms)
  include_ext = ["go", "html"] # Watched extensions
  exclude_dir = ["vendor"]   # Ignored directories
```

For full configuration options, see [Air documentation](https://github.com/cosmtrek/air).
