# AGENTS.md - Development Guide for SoftFetch

## Project Overview
SoftFetch is a Go project for fetching software information. The project follows standard Go project layout with clear separation between executable and library code.

## Build & Test Commands

### Standard Commands (via Make)
- `make build` - Build the application to `bin/softfetch`
- `make test` - Run all tests with verbose output
- `make test-coverage` - Run tests with coverage report
- `make lint` - Run golangci-lint (install with `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`)
- `make fmt` - Format all code with `go fmt`
- `make clean` - Remove build artifacts
- `make deps` - Install dependencies and tidy modules
- `make run` - Build and run the application

### Direct Go Commands
- `go build -o bin/softfetch ./cmd/softfetch` - Build without Make
- `go test ./...` - Run all tests
- `go test -v ./pkg/version/` - Run tests in specific package
- `go test -run TestGetVersion ./pkg/version/` - Run single test by name
- `go test -v -run "TestGet.*" ./pkg/version/` - Run tests matching pattern

### Single Test Execution
```bash
# Run a specific test function
go test -v -run TestFunctionName ./path/to/package

# Run tests matching a pattern
go test -v -run ".*Pattern.*" ./...

# Run with timeout
go test -v -timeout 30s ./...
```

## Code Style Guidelines

### Import Organization
1. Standard library imports first
2. Third-party imports second (separated by blank line)
3. Local project imports third (separated by blank line)
```go
import (
    "fmt"
    "os"
    
    "github.com/spf13/cobra"
    
    "github.com/gausszhou/softfetch/pkg/version"
)
```

### Formatting
- Use `go fmt` or `gofmt` for all formatting
- Tabs for indentation, not spaces
- Maximum line length: 120 characters (soft limit)
- No unused imports or variables

### Naming Conventions
- **Packages**: Short, single-word, lowercase (e.g., `version`, `config`)
- **Functions**: CamelCase for exported, camelCase for unexported
- **Variables**: camelCase for local, PascalCase for exported
- **Constants**: PascalCase for exported, camelCase for unexported
- **Interfaces**: Descriptive names, often with `er` suffix (e.g., `Reader`, `Writer`)
- **Test functions**: `TestFunctionName` in `*_test.go` files

### Type Guidelines
- Use specific types over generic ones (e.g., `int64` over `int` when size matters)
- Prefer structs for complex data, interfaces for behavior
- Use type aliases for domain-specific types
- Document all exported types

### Error Handling
- Always check and handle errors explicitly
- Use `fmt.Errorf("context: %w", err)` for wrapping errors
- Define custom error types for specific error conditions
- Return errors as last return value
- Log errors with context, not just the error message

### Documentation
- All exported functions, types, and variables must have comments
- Comments should start with the name of the thing being documented
- Keep comments concise and informative
- Use `go doc` style comments

### Testing
- Test file naming: `*_test.go` in same package
- Table-driven tests for multiple test cases
- Use `t.Helper()` in helper functions
- Benchmark functions: `BenchmarkFunctionName`
- Test both success and error paths

## Project Structure
```
main.go           # Main application entry point
pkg/version/      # Exported version package
internal/         # Private application code
configs/          # Configuration files
scripts/          # Build and CI scripts
deployments/      # Deployment configurations
docs/             # Documentation
test/             # Integration and end-to-end tests
```

## Git Workflow
- Commit messages: imperative mood, present tense
- Keep commits focused and atomic
- Run `make fmt` and `make test` before committing
- Feature branches for new development

## Dependencies
- Use `go mod tidy` to clean dependencies
- Minimize external dependencies
- Use `go get` to add new dependencies
- Check licenses for compatibility