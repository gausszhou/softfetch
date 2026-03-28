package display

import (
	"fmt"
	"strings"

	"github.com/gausszhou/softfetch/internal/detect"
)

var (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
	ColorGray   = "\033[90m"
)

func PrintResult(result detect.DetectionResult) {
	fmt.Println()
	fmt.Println(ColorCyan + "╭──────────────────────────────────────────────────────╮" + ColorReset)
	fmt.Println(ColorCyan + "│" + ColorReset + "           " + ColorBlue + "SoftFetch" + ColorReset + " - System Information Tool        " + ColorCyan + "│" + ColorReset)
	fmt.Println(ColorCyan + "╰──────────────────────────────────────────────────────╯" + ColorReset)
	fmt.Println()

	maxNameLen := 0
	maxVerLen := 0
	for _, tool := range result.Tools {
		if len(tool.Name) > maxNameLen {
			maxNameLen = len(tool.Name)
		}
		if len(tool.Version) > maxVerLen {
			maxVerLen = len(tool.Version)
		}
	}

	nameWidth := maxNameLen
	if nameWidth < 12 {
		nameWidth = 12
	}
	verWidth := maxVerLen
	if verWidth < 20 {
		verWidth = 20
	}

	fmt.Printf(ColorGray+"  %-"+fmt.Sprintf("%d", nameWidth)+"s  %-"+fmt.Sprintf("%d", verWidth)+"s  %s"+ColorReset+"\n",
		"Environment", "Version", "Path")
	fmt.Println(strings.Repeat("─", nameWidth+verWidth+40))

	for _, tool := range result.Tools {
		printTool(tool, nameWidth, verWidth)
	}

	fmt.Println()
	fmt.Println(ColorGray + "  " + result.OS + "/" + result.Arch + ColorReset)
	fmt.Println()
}

func printTool(tool detect.Tool, nameWidth, verWidth int) {
	name := tool.Name
	name = name + strings.Repeat(" ", nameWidth-len(name))

	version := tool.Version
	path := tool.Path

	var versionColor string
	if tool.Detected {
		versionColor = ColorGreen
	} else {
		versionColor = ColorRed
	}

	fmt.Printf("  %s%-14s%s  %s%-20s%s  %s\n",
		ColorYellow,
		name,
		ColorReset,
		versionColor,
		version,
		ColorReset,
		ColorGray+path+ColorReset,
	)
}

func PrintSimple(result detect.DetectionResult) {
	for _, tool := range result.Tools {
		if tool.Detected {
			fmt.Printf("%-12s %-20s (%s)\n", tool.Name+":", tool.Version, tool.Path)
		} else {
			fmt.Printf("%-12s %s\n", tool.Name+":", ColorRed+"Not installed"+ColorReset)
		}
	}
}
