package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Init initializes Viper configuration with defaults, config file, and env vars
func Init(cfgFile string) error {
	// Set defaults
	viper.SetDefault("app.name", "gotemplate-cli")
	viper.SetDefault("app.version", "0.1.0")
	viper.SetDefault("app.environment", "development")
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.format", "text")
	viper.SetDefault("output.color", true)
	viper.SetDefault("output.quiet", false)

	// Environment variables
	viper.SetEnvPrefix("GOTEMPLATE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Config file
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.gotemplate-cli")
		viper.AddConfigPath("/etc/gotemplate-cli")
	}

	// Read config file (ok if not found)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("error reading config file: %w", err)
		}
		// Config file not found; use defaults + env vars
	}

	return nil
}

// GetConfigFilePath returns the config file path being used
func GetConfigFilePath() string {
	return viper.ConfigFileUsed()
}

// WriteDefaultConfig writes a default config file to the given path
func WriteDefaultConfig(path string) error {
	defaultConfig := `# gotemplate-cli configuration
app:
  name: gotemplate-cli
  version: 0.1.0
  environment: development

log:
  level: info    # debug, info, warn, error
  format: text   # text, json

output:
  color: true
  quiet: false
`
	return os.WriteFile(path, []byte(defaultConfig), 0644)
}
