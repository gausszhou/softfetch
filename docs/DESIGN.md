# Design

## Project Structure

```
.
├── main.go              # Application entry point
├── pkg/                 # Exported packages
│   └── version/         # Version information
├── internal/            # Private application code
│   ├── command/         # Command execution utilities
│   ├── detect/         # Tool detection logic
│   └── display/        # Output formatting
├── configs/             # Configuration files
├── scripts/             # Build and CI scripts
├── deployments/         # Deployment configurations
├── docs/                # Documentation
└── test/               # Integration tests
```

## Architecture

### Detection System

The detection system uses a detector interface to identify installed tools:

- Each detector implements the `Detector` interface
- `GetCoreDetectors()` returns all supported detectors
- `Detect()` runs all detectors and aggregates results

### Command Execution

The `command` package provides utilities for executing external commands with timeout support.
