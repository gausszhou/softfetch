package detect

import (
	"testing"
)

func TestGoDetector(t *testing.T) {
	detector := NewGoDetector()
	tool := detector.Detect()

	if detector.Name() != "Go" {
		t.Errorf("Expected name 'Go', got '%s'", detector.Name())
	}

	if detector.Category() != CategoryLanguage {
		t.Errorf("Expected category 'Language', got '%s'", detector.Category())
	}

	if tool.Name != "Go" {
		t.Errorf("Expected tool name 'Go', got '%s'", tool.Name)
	}
}

func TestNodeDetector(t *testing.T) {
	detector := NewNodeDetector()
	tool := detector.Detect()

	if detector.Name() != "Node.js" {
		t.Errorf("Expected name 'Node.js', got '%s'", detector.Name())
	}

	if tool.Name != "Node.js" {
		t.Errorf("Expected tool name 'Node.js', got '%s'", tool.Name)
	}
}

func TestPythonDetector(t *testing.T) {
	detector := NewPythonDetector()
	tool := detector.Detect()

	if detector.Name() != "Python" {
		t.Errorf("Expected name 'Python', got '%s'", detector.Name())
	}

	if tool.Name != "Python" {
		t.Errorf("Expected tool name 'Python', got '%s'", tool.Name)
	}
}

func TestJavaDetector(t *testing.T) {
	detector := NewJavaDetector()
	tool := detector.Detect()

	if detector.Name() != "Java" {
		t.Errorf("Expected name 'Java', got '%s'", detector.Name())
	}

	if tool.Name != "Java" {
		t.Errorf("Expected tool name 'Java', got '%s'", tool.Name)
	}
}

func TestGetCoreDetectors(t *testing.T) {
	detectors := GetCoreDetectors()

	if len(detectors) != 4 {
		t.Errorf("Expected 4 detectors, got %d", len(detectors))
	}

	expectedNames := []string{"Go", "Node.js", "Python", "Java"}
	for i, d := range detectors {
		if d.Name() != expectedNames[i] {
			t.Errorf("Expected detector[%d] name '%s', got '%s'", i, expectedNames[i], d.Name())
		}
	}
}

func TestDetect(t *testing.T) {
	detectors := GetCoreDetectors()
	result := Detect(detectors...)

	if result.OS == "" {
		t.Error("Expected OS to be set")
	}

	if result.Arch == "" {
		t.Error("Expected Arch to be set")
	}

	if len(result.Tools) != 4 {
		t.Errorf("Expected 4 tools in result, got %d", len(result.Tools))
	}
}

func TestParseVersion(t *testing.T) {
	tests := []struct {
		name     string
		output   string
		expected string
	}{
		{"Go", "go version go1.21.0 darwin/amd64", "go1.21.0"},
		{"Node", "v20.10.0", "v20.10.0"},
		{"Python", "Python 3.11.6", "3.11.6"},
		{"Java", "openjdk version \"21.0.1\" 2023-10-17", "21.0.1"},
		{"Empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseVersion(tt.output)
			if tt.expected != "" && result == "" {
				t.Errorf("Expected version '%s', got empty string", tt.expected)
			}
		})
	}
}
