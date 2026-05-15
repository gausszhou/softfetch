package cmd

import (
	"github.com/gausszhou/softfetch/internal/detect"
	"github.com/gausszhou/softfetch/internal/display"
	"github.com/gausszhou/softfetch/internal/info"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "softfetch",
	Short: "A command-line tool that detects installed development tools and languages",
	Long:  `SoftFetch is a simple and fast CLI tool to detect installed development tools, languages, and system information.`,
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag, err := cmd.Flags().GetBool("version")
		if err == nil && versionFlag {
			display.PrintVersion(info.Version)
			return
		}
		detectors := detect.GetCoreDetectors()
		result := detect.Detect(detectors...)
		display.PrintResult(result)
	},
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print version information")
}

func Execute() error {
	return rootCmd.Execute()
}
