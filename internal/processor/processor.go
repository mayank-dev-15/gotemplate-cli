package processor

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Processor handles the main processing logic
type Processor struct {
	StartedAt time.Time
}

// New creates a new Processor instance
func New() *Processor {
	return &Processor{
		StartedAt: time.Now(),
	}
}

// Process reads the input file, processes the content, and writes output
func (p *Processor) Process(inputPath, outputPath string) (string, error) {
	// Read input
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return "", fmt.Errorf("failed to read input file: %w", err)
	}

	content := string(data)

	// Process the content
	result := p.transform(content)

	// Write output if specified
	if outputPath != "" {
		if err := os.WriteFile(outputPath, []byte(result), 0644); err != nil {
			return "", fmt.Errorf("failed to write output file: %w", err)
		}
	}

	return fmt.Sprintf("processed %d bytes in %v", len(data), time.Since(p.StartedAt)), nil
}

// DryRun simulates processing without writing output
func (p *Processor) DryRun(inputPath string) (string, error) {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return "", fmt.Errorf("failed to read input file: %w", err)
	}

	return fmt.Sprintf("would process %d bytes from %s", len(data), inputPath), nil
}

// transform applies processing transformations to the content
func (p *Processor) transform(content string) string {
	// Example transformations
	result := content
	result = strings.TrimSpace(result)
	result = strings.ToUpper(result)
	return result
}
