# SoftFetch

A Go project for fetching software information.

## Project Structure

```
.
├── cmd/softfetch/      # Application entry points
├── pkg/               # Library code that's safe to export
├── internal/          # Private application and library code
├── configs/           # Configuration files
├── scripts/           # Scripts for building, CI, etc.
├── deployments/       # Deployment configurations
├── docs/              # Documentation
├── test/              # Additional test files
└── vendor/            # Vendor dependencies (if used)
```

## Getting Started

### Prerequisites

- Go 1.21 or later

### Installation

```bash
# Clone the repository
git clone https://github.com/user/softfetch.git
cd softfetch

# Install dependencies
go mod download

# Build the application
go build -o bin/softfetch ./cmd/softfetch

# Run the application
./bin/softfetch
```

## Development

### Build

```bash
make build
```

### Test

```bash
make test
```

### Lint

```bash
make lint
```

## License

MIT License