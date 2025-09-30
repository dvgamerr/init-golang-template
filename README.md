# Go Project Template 2025

A modern Go project template following the latest best practices and standard structure.

## Project Structure

```
.
├── cmd/                    # Main applications
│   └── app/               # Application entry point
├── internal/              # Private application code
│   ├── config/           # Configuration
│   ├── handler/          # HTTP handlers
│   ├── service/          # Business logic
│   └── repository/       # Data access layer
├── pkg/                   # Public libraries
├── api/                   # API definitions (OpenAPI/Swagger)
├── configs/               # Configuration files
├── scripts/               # Scripts for build, install, etc.
├── test/                  # Additional test data
├── docs/                  # Documentation
└── .github/              # GitHub specific files
    └── workflows/        # CI/CD workflows
```

## Prerequisites

- Go 1.23 or higher
- golangci-lint for linting
- commitlint for commit message validation

## Installation

```bash
# Install dependencies
go mod download

# Install development tools
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

## Development

```bash
# Run the application
go run cmd/app/main.go

# Build
go build -o bin/app cmd/app/main.go

# Run tests
go test ./...

# Run linter
golangci-lint run

# Run tests with coverage
go test -v -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

## Linting

This project uses `golangci-lint` with a comprehensive configuration.

```bash
# Run all linters
golangci-lint run

# Run with auto-fix
golangci-lint run --fix
```

## Commit Convention

This project uses **JIRA ID + Conventional Commits** format.

Format: `JIRA-123/type(scope): subject`

The JIRA ID is optional but recommended for tracking.

### Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `perf`: Performance improvements
- `test`: Adding or updating tests
- `build`: Build system or dependencies
- `ci`: CI configuration
- `chore`: Maintenance tasks
- `revert`: Revert previous commit

### Examples:
```bash
# With JIRA ID
git commit -m "JIRA-123/feat: add user authentication"
git commit -m "JIRA-456/fix(api): resolve timeout issue"
git commit -m "PROJ-789/docs: update API documentation"

# Without JIRA ID (also valid)
git commit -m "feat(auth): implement JWT validation"
git commit -m "fix: resolve memory leak"
```

📖 See [COMMIT_GUIDELINES.md](docs/COMMIT_GUIDELINES.md) for detailed documentation.

## License

MIT
