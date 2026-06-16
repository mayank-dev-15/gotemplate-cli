# gotemplate-cli

[![Live Demo](https://img.shields.io/badge/🚀_Live_Demo-Visit-blue?style=for-the-badge)](https://mayank-dev-15.github.io/gotemplate-cli-demo)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Language](https://img.shields.io/badge/Language-Go-green)

Professional Go CLI template with cobra, viper, logrus, and testing setup.

`go` `cli` `template` `cobra`

---

## ✨ Features

- cobra command structure with subcommands
- viper config management (YAML, JSON, ENV)
- logrus structured logging
- Unit tests with testing package
- Shell completion generation
- Man page generation
- Cross-compilation support

---

## 🚀 Live Demo

**[View Demo →](https://mayank-dev-15.github.io/gotemplate-cli-demo)**

The demo is hosted on GitHub Pages. No installation needed — just click and explore.

---

## 🛠️ Tech Stack

- Go 1.21+
- cobra
- viper
- logrus

---

## 📦 Installation

```bash
git clone https://github.com/mayank-dev-15/gotemplate-cli.git
cd gotemplate-cli
```

```bash
cd gotemplate-cli
go mod tidy
go build -o gotemplate .
./gottemplate --help
```

---

## 💡 Usage

```bash
# Build and run
go run main.go --help

# With config
go run main.go --config config.yaml serve

# Run tests
go test ./...

# Cross-compile
GOOS=linux GOARCH=amd64 go build -o gotemplate-linux
```

---

## 📁 Project Structure

```
gotemplate-cli/
├── README.md          # This file
├── Demo.md            # Demo documentation
├── LICENSE            # MIT License
└── ...                # Source files
```

---

## 🤝 Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

---

## 📄 License

This project is licensed under the MIT License.

---

## 🔗 Links

- **Live Demo:** [https://mayank-dev-15.github.io/gotemplate-cli-demo](https://mayank-dev-15.github.io/gotemplate-cli-demo)
- **Source Code:** [github.com/mayank-dev-15/gotemplate-cli](https://github.com/mayank-dev-15/gotemplate-cli)
- **Issues:** [github.com/mayank-dev-15/gotemplate-cli/issues](https://github.com/mayank-dev-15/gotemplate-cli/issues)
- **Releases:** [github.com/mayank-dev-15/gotemplate-cli/releases](https://github.com/mayank-dev-15/gotemplate-cli/releases)
- **Demo Docs:** [Demo.md](https://github.com/mayank-dev-15/gotemplate-cli/blob/main/Demo.md)

---

*Built with ❤️ by [Mayank Basena](https://github.com/mayank-dev-15) · 15 · GSoC 2027 Aspirant*

---

## ⚠️ Attribution & Credit Notice

This project is created and maintained by **Mayank Basena** ([@mayank-dev-15](https://github.com/mayank-dev-15)).

If you fork, use, modify, or derive work from this repository, **you must give proper credit** to the original author. This includes:

- Keeping this attribution section intact in any fork or derivative work
- Crediting **Mayank Basena** in your project's README or documentation
- Linking back to the original repository

**Failure to provide proper credit is a violation of the spirit of open source and may result in a DMCA takedown request.**

> *"No AI. No Shortcuts."* — Mayank Basena
