# SoftFetch

A command-line tool that detects installed development tools and languages on your system.

## Features

- Detects: Go, Node.js, Python, Java, C, C++, Rust, PHP
- Displays OS and architecture information
- Simple and fast

## Installation

```bash
# Install latest version
go install github.com/gausszhou/softfetch@latest

# Or build from source
git clone https://github.com/gausszhou/softfetch.git
cd softfetch
make install
```

## Usage

```bash
# Run the tool
softfetch

# Check version
softfetch --version
```

## Clean

```bash
# Clean build artifacts
make clean

# Clean Go module cache (all cached versions)
go clean -cache

# Remove installed binary
rm -f $(go env GOPATH)/bin/softfetch
```

## License

MIT
