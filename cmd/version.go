package cmd

import (
	"runtime"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Display the version, build info, and Go runtime details.`,
	Run: func(cmd *cobra.Command, args []string) {
		green := color.New(color.FgGreen, color.Bold)
		cyan := color.New(color.FgCyan)

		green.Println("gotemplate-cli")
		cyan.Printf("  Version:    %s\n", version)
		cyan.Printf("  Go version: %s\n", runtime.Version())
		cyan.Printf("  OS/Arch:    %s/%s\n", runtime.GOOS, runtime.GOARCH)
		cyan.Printf("  Compiler:   %s\n", runtime.Compiler)
	},
}
