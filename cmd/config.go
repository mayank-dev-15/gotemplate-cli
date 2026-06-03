package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/mayank-dev-15/gotemplate-cli/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configKey   string
	configValue string
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View or modify configuration",
	Long: `Display or edit configuration values stored in the config file.
Configuration is managed via Viper and supports YAML files and environment variables.

Environment variables are automatically mapped with the prefix GOTEMPLATE.
For example, GOTEMPLATE_LOG_LEVEL overrides the log.level config key.`,
	Example: `  gotemplate-cli config
  gotemplate-cli config --key log.level
  gotemplate-cli config --key log.level --value debug`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Show specific key
		if configKey != "" {
			if configValue != "" {
				// Set the value
				viper.Set(configKey, configValue)
				if err := viper.WriteConfig(); err != nil {
					return fmt.Errorf("failed to write config: %w", err)
				}
				green := color.New(color.FgGreen, color.Bold)
				green.Printf("✓ Set %s = %s\n", configKey, configValue)
				logger.Log.WithFields(map[string]interface{}{
					"key":   configKey,
					"value": configValue,
				}).Info("Configuration updated")
				return nil
			}

			// Get the value
			val := viper.Get(configKey)
			if val == nil {
				yellow := color.New(color.FgYellow)
				yellow.Printf("Key '%s' is not set\n", configKey)
				return nil
			}

			cyan := color.New(color.FgCyan)
			cyan.Printf("%s = %v\n", configKey, val)
			return nil
		}

		// Show all config
		cyan := color.New(color.FgCyan, color.Bold)
		cyan.Println("Current configuration:")
		cyan.Println("─────────────────────")

		settings := viper.AllSettings()
		for key, val := range settings {
			white := color.New(color.FgWhite)
			white.Printf("  %-20s = %v\n", key, val)
		}

		configFile := viper.ConfigFileUsed()
		if configFile != "" {
			cyan.Println("─────────────────────")
			cyan.Printf("Config file: %s\n", configFile)
		}

		logger.Log.Info("Displayed configuration")
		return nil
	},
}

func init() {
	configCmd.Flags().StringVarP(&configKey, "key", "k", "", "Configuration key to get/set")
	configCmd.Flags().StringVarP(&configValue, "value", "V", "", "Value to set (if setting a key)")
}
