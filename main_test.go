package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/gausszhou/softfetch/internal/version"
)

func TestVersionFlag(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"versionLong", []string{"--version"}},
		{"versionShort", []string{"-v"}},
	}

	expectedVersion := version.GetVersion()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "."}, tt.args...)...)
			cmd.Env = append(os.Environ(), "TERM=dumb")
			output, err := cmd.CombinedOutput()

			if err != nil {
				t.Fatalf("Command failed: %v, output: %s", err, output)
			}

			expectedOutput := "SoftFetch " + expectedVersion
			if !strings.Contains(string(output), expectedOutput) {
				t.Errorf("Expected output to contain '%s', got '%s'", expectedOutput, string(output))
			}
		})
	}
}

func TestMainExecution(t *testing.T) {
	cmd := exec.Command("go", "run", ".")
	cmd.Env = append(os.Environ(), "TERM=dumb")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("Command failed: %v, output: %s", err, output)
	}

	if len(output) == 0 {
		t.Error("Expected output from main execution")
	}
}
