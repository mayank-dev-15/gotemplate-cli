package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/mayank-dev-15/gotemplate-cli/internal/logger"
	"github.com/mayank-dev-15/gotemplate-cli/internal/processor"
	"github.com/spf13/cobra"
)

var (
	inputFile string
	outputFile string
	dryRun    bool
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the main processing pipeline",
	Long: `Execute the main processing pipeline with the given input.
This command reads input, processes it through the pipeline,
and produces output according to the configuration.`,
	Example: `  gotemplate-cli run --input data.txt
  gotemplate-cli run --input data.txt --output result.txt
  gotemplate-cli run --input data.txt --dry-run`,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.Log.Info("Starting processing pipeline")

		// Validate input
		if inputFile == "" && len(args) == 0 {
			return fmt.Errorf("input file is required (use --input or provide as positional argument)")
		}

		// Positional argument takes precedence
		if len(args) > 0 {
			inputFile = args[0]
		}

		// Validate the input
		if err := processor.ValidateInput(inputFile); err != nil {
			return fmt.Errorf("validation failed: %w", err)
		}

		// Create processor
		p := processor.New()

		if dryRun {
			yellow := color.New(color.FgYellow, color.Bold)
			yellow.Println("⚡ Dry-run mode — no changes will be made")
			logger.Log.Info("Dry-run mode enabled")

			result, err := p.DryRun(inputFile)
			if err != nil {
				return fmt.Errorf("dry-run failed: %w", err)
			}

			cyan := color.New(color.FgCyan)
			cyan.Printf("Would process: %s\n", result)
			logger.Log.WithField("result", result).Info("Dry-run completed")
			return nil
		}

		// Run the processor
		result, err := p.Process(inputFile, outputFile)
		if err != nil {
			return fmt.Errorf("processing failed: %w", err)
		}

		green := color.New(color.FgGreen, color.Bold)
		green.Println("✓ Processing complete!")
		cyan := color.New(color.FgCyan)
		cyan.Printf("  Result: %s\n", result)

		if outputFile != "" {
			cyan.Printf("  Output written to: %s\n", outputFile)
		}

		logger.Log.WithFields(map[string]interface{}{
			"input":  inputFile,
			"output": outputFile,
			"result": result,
		}).Info("Processing completed successfully")

		return nil
	},
}

func init() {
	runCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file path")
	runCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file path (default: stdout)")
	runCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Simulate processing without making changes")
}
