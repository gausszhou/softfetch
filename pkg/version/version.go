package version

import "fmt"

var (
	// Version is the current version of the application
	Version = "0.1.0"

	// BuildDate is the date when the binary was built
	BuildDate = "unknown"

	// GitCommit is the git commit hash
	GitCommit = "unknown"
)

// GetVersion returns the version string
func GetVersion() string {
	return Version
}

// GetBuildInfo returns build information
func GetBuildInfo() string {
	return fmt.Sprintf("Version: %s\nBuild Date: %s\nGit Commit: %s", Version, BuildDate, GitCommit)
}
