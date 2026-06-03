package cmd

import (
	"fmt"
	"os"

	"github.com/mayank-dev-15/gotemplate-cli/internal/config"
	"github.com/mayank-dev-15/gotemplate-cli/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	verbose bool
	version = "0.1.0"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gotemplate-cli",
	Short: "A feature-rich Go CLI template with cobra, viper, and logrus",
	Long: `gotemplate-cli is a production-ready CLI tool template built with:
  • Cobra for CLI commands and subcommands
  • Viper for configuration management (YAML + env vars)
  • Logrus for structured logging
  • Fatih/color for colored output

Use it as a starting point for building your own Go CLI tools.`,
	Version: version,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Skip config loading for version command
		if cmd.Name() == "version" {
			return nil
		}
		if err := config.Init(cfgFile); err != nil {
			return fmt.Errorf("failed to init config: %w", err)
		}
		logger.Init(verbose)
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gotemplate-cli.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Bind flags to viper
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	// Add subcommands
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(versionCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
