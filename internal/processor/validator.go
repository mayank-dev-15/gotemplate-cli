package processor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ValidateInput checks if the input file is valid for processing
func ValidateInput(inputPath string) error {
	if inputPath == "" {
		return fmt.Errorf("input path cannot be empty")
	}

	// Check file exists
	info, err := os.Stat(inputPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("input file does not exist: %s", inputPath)
		}
		return fmt.Errorf("failed to stat input file: %w", err)
	}

	// Check it's not a directory
	if info.IsDir() {
		return fmt.Errorf("input path is a directory, not a file: %s", inputPath)
	}

	// Check file is not empty
	if info.Size() == 0 {
		return fmt.Errorf("input file is empty: %s", inputPath)
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(inputPath))
	allowedExts := map[string]bool{
		".txt":  true,
		".md":   true,
		".json": true,
		".yaml": true,
		".yml":  true,
		".csv":  true,
		".log":  true,
	}

	if !allowedExts[ext] {
		return fmt.Errorf("unsupported file type '%s' — allowed: .txt, .md, .json, .yaml, .yml, .csv, .log", ext)
	}

	return nil
}
