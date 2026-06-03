package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/mayank-dev-15/gotemplate-cli/internal/logger"
	"github.com/spf13/cobra"
)

var (
	projectName string
	projectDir  string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project scaffold",
	Long: `Create a new project scaffold with a standard directory structure
and configuration files. This sets up everything you need to get started.`,
	Example: `  gotemplate-cli init --name myproject
  gotemplate-cli init --name myproject --dir /path/to/create`,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.Log.WithField("project", projectName).Info("Initializing new project")

		// Determine target directory
		targetDir := projectDir
		if targetDir == "" {
			targetDir = projectName
		}

		// Check if directory already exists
		if _, err := os.Stat(targetDir); !os.IsNotExist(err) {
			return fmt.Errorf("directory %s already exists", targetDir)
		}

		// Create directory structure
		dirs := []string{
			targetDir,
			filepath.Join(targetDir, "cmd"),
			filepath.Join(targetDir, "internal"),
			filepath.Join(targetDir, "pkg"),
			filepath.Join(targetDir, "configs"),
			filepath.Join(targetDir, "scripts"),
		}

		for _, dir := range dirs {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", dir, err)
			}
		}

		// Create a default config file
		configContent := fmt.Sprintf(`# %s configuration
app:
  name: %s
  version: 0.1.0
  environment: development

log:
  level: info
  format: text

output:
  color: true
  quiet: false
`, projectName, projectName)

		configPath := filepath.Join(targetDir, "configs", "config.yaml")
		if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
			return fmt.Errorf("failed to write config file: %w", err)
		}

		// Create a .gitignore
		gitignoreContent := `# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary
*.test

# Output
/out/
/dist/
/build/

# Dependencies
/vendor/

# IDE
.idea/
.vscode/
*.swp
*.swo

# OS
.DS_Store
Thumbs.db

# Config local overrides
config.local.yaml
`
		gitignorePath := filepath.Join(targetDir, ".gitignore")
		if err := os.WriteFile(gitignorePath, []byte(gitignoreContent), 0644); err != nil {
			return fmt.Errorf("failed to write .gitignore: %w", err)
		}

		// Create a README stub
		readmeContent := fmt.Sprintf(`# %s

## Getting Started

`+"```bash"+`
go build -o %s .
./%s --help
`+"```"+`

## Configuration

Edit %s to customize settings.

## License
MIT
`, projectName, projectName, projectName, configPath)
		readmePath := filepath.Join(targetDir, "README.md")
		if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
			return fmt.Errorf("failed to write README: %w", err)
		}

		green := color.New(color.FgGreen, color.Bold)
		cyan := color.New(color.FgCyan)

		green.Println("✓ Project scaffold created successfully!")
		cyan.Printf("  Project: %s\n", projectName)
		cyan.Printf("  Location: %s\n", targetDir)
		cyan.Println("\nNext steps:")
		cyan.Printf("  cd %s\n", targetDir)
		cyan.Println("  go mod init <module-path>")
		cyan.Println("  go build -o", projectName)

		logger.Log.WithField("project", projectName).Info("Project initialized successfully")
		return nil
	},
}

func init() {
	initCmd.Flags().StringVarP(&projectName, "name", "n", "", "Name of the project (required)")
	initCmd.Flags().StringVarP(&projectDir, "dir", "d", "", "Directory to create (defaults to project name)")
	initCmd.MarkFlagRequired("name")
}
