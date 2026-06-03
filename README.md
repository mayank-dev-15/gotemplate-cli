# GoTemplate CLI

[![Go Version](https://img.shields.io/badge/go-1.22+-00ADD8?logo=go&logoColor=white)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Build](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://github.com/mayank-dev-15/gotemplate-cli)
[![Release](https://img.shields.io/badge/release-v0.1.0-blue.svg)](https://github.com/mayank-dev-15/gotemplate-cli/releases)

A production-ready CLI tool template built with **Go**, featuring command subcommands, configuration management, structured logging, and colored output. Use this as a starting point for building your own Go CLI tools.

## ✨ Features

- **Cobra** — Elegant command structure with subcommands and flag management
- **Viper** — Hierarchical config via YAML files, environment variables, and flags
- **Logrus** — Structured logging with configurable levels (debug, info, warn, error)
- **Fatih/color** — Beautiful colored terminal output
- **Project scaffolding** — `init` command creates new project skeletons
- **Dry-run mode** — Simulate processing without side effects
- **Input validation** — File existence, size, and extension checking

## 📦 Installation

### Build from source

```bash
git clone https://github.com/mayank-dev-15/gotemplate-cli.git
cd gotemplate-cli
go build -o gotemplate-cli .
```

### Install directly

```bash
go install github.com/mayank-dev-15/gotemplate-cli@latest
```

## 🚀 Quick Start

```bash
# Show help
./gotemplate-cli --help

# Show version
./gotemplate-cli version

# Initialize a new project scaffold
./gotemplate-cli init --name myproject

# Run the processing pipeline
./gotemplate-cli run --input sample.txt

# Run with output file and verbose logging
./gotemplate-cli run --input data.txt --output result.txt --verbose

# Dry-run mode (no writes)
./gotemplate-cli run --input data.txt --dry-run

# View configuration
./gotemplate-cli config

# Set a config value
./gotemplate-cli config --key log.level --value debug
```

## 🧰 Commands

| Command   | Description                                    | Key Flags                                    |
|-----------|------------------------------------------------|----------------------------------------------|
| `root`    | Base command — shows help and global flags      | `--config`, `--verbose`                      |
| `init`    | Scaffold a new project with sample structure    | `--name`, `--dir`                            |
| `run`     | Execute the processing pipeline                 | `--input`, `--output`, `--dry-run`           |
| `version` | Print version and Go runtime info               | —                                            |
| `config`  | View or modify configuration values             | `--key`, `--value`                           |

## ⚙️ Configuration

Configuration is managed via **Viper** with the following priority order:

1. **Flags** — Command-line flags take highest precedence
2. **Environment variables** — Prefixed with `GOTEMPLATE_` (e.g., `GOTEMPLATE_LOG_LEVEL=debug`)
3. **Config file** — YAML file at `./config.yaml`, `$HOME/.gotemplate-cli.yaml`, or `/etc/gotemplate-cli/config.yaml`
4. **Defaults** — Built-in sensible defaults

### Example `config.yaml`

```yaml
app:
  name: gotemplate-cli
  version: 0.1.0
  environment: development

log:
  level: info       # debug | info | warn | error
  format: text      # text | json

output:
  color: true
  quiet: false
```

### Supported input file types

The `run` command validates input files against these extensions:

`.txt` · `.md` · `.json` · `.yaml` · `.yml` · `.csv` · `.log`

## 🛠️ Build Targets

```bash
make build      # Compile the binary
make test       # Run all tests
make testv      # Tests with verbose output and coverage
make run        # Build and show help
make clean      # Remove artifacts
make lint       # Run go vet and gofmt
make deps       # Download and tidy modules
make install    # Install to GOPATH/bin
make fmt        # Format all Go files
make all        # Build and test
```

## 📁 Project Structure

```
gotemplate-cli/
├── main.go                    # Entry point
├── go.mod                     # Go module definition
├── Makefile                   # Build automation
├── README.md                  # This file
├── .gitignore                 # Git ignore rules
├── cmd/
│   ├── root.go                # Root command + global flags
│   ├── run.go                 #  run subcommand
│   ├── init.go                #  init subcommand
│   ├── version.go             #  version subcommand
│   └── config.go              #  config subcommand
└── internal/
    ├── config/
    │   └── config.go          # Viper config initialization
    ├── logger/
    │   └── logger.go          # Logrus logger setup
    └── processor/
        ├── processor.go       # Core processing logic
        └── validator.go       # Input validation
```

## 🧪 Testing

```bash
go test ./...
go test -v -cover ./...
```

## 📝 License

[MIT](LICENSE) — feel free to use, modify, and distribute.

---

Built with ❤️ by [Mayank Basena](https://github.com/mayank-dev-15)
