package version

import (
	"testing"
)

func TestGetVersion(t *testing.T) {
	version := GetVersion()
	if version == "" {
		t.Error("GetVersion() returned empty string")
	}
	if version != Version {
		t.Errorf("GetVersion() = %v, want %v", version, Version)
	}
}

func TestGetBuildInfo(t *testing.T) {
	buildInfo := GetBuildInfo()
	if buildInfo == "" {
		t.Error("GetBuildInfo() returned empty string")
	}
	// Check that build info contains version string
	if !contains(buildInfo, Version) {
		t.Errorf("GetBuildInfo() does not contain version %v", Version)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || containsMiddle(s, substr)))
}

func containsMiddle(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
