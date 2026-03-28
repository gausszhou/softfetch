package detect

import (
	"github.com/user/softfetch/internal/command"
)

type GoDetector struct {
	baseDetector
}

func NewGoDetector() *GoDetector {
	return &GoDetector{
		baseDetector: baseDetector{
			name:     "Go",
			category: CategoryLanguage,
		},
	}
}

func (d *GoDetector) Detect() Tool {
	result := command.Execute("go", "version")
	tool := Tool{
		Name:     d.name,
		Category: string(d.Category()),
		Symbol:   "⬢",
	}

	if !result.Exists {
		tool.Detected = false
		tool.Version = "Not installed"
		return tool
	}

	if result.Error != nil {
		tool.Detected = false
		tool.Version = parseVersion(result.Output)
		return tool
	}

	path, _ := command.LookPath("go")
	tool.Path = path
	tool.Detected = true
	tool.Version = parseVersion(result.Output)
	return tool
}

type NodeDetector struct {
	baseDetector
}

func NewNodeDetector() *NodeDetector {
	return &NodeDetector{
		baseDetector: baseDetector{
			name:     "Node.js",
			category: CategoryLanguage,
		},
	}
}

func (d *NodeDetector) Detect() Tool {
	result := command.Execute("node", "--version")
	tool := Tool{
		Name:     d.name,
		Category: string(d.Category()),
		Symbol:   "⬢",
	}

	if !result.Exists {
		tool.Detected = false
		tool.Version = "Not installed"
		return tool
	}

	if result.Error != nil {
		tool.Detected = false
		tool.Version = parseVersion(result.Output)
		return tool
	}

	path, _ := command.LookPath("node")
	tool.Path = path
	tool.Detected = true
	tool.Version = parseVersion(result.Output)
	return tool
}

type PythonDetector struct {
	baseDetector
}

func NewPythonDetector() *PythonDetector {
	return &PythonDetector{
		baseDetector: baseDetector{
			name:     "Python",
			category: CategoryLanguage,
		},
	}
}

func (d *PythonDetector) Detect() Tool {
	tool := Tool{
		Name:     d.name,
		Category: string(d.Category()),
		Symbol:   "🐍",
	}

	commands := []string{"python3", "python"}
	for _, cmd := range commands {
		result := command.Execute(cmd, "--version")
		if result.Exists && result.Error == nil {
			tool.Detected = true
			tool.Version = parseVersion(result.Output)
			path, _ := command.LookPath(cmd)
			tool.Path = path
			return tool
		}
	}

	tool.Detected = false
	tool.Version = "Not installed"
	return tool
}

type JavaDetector struct {
	baseDetector
}

func NewJavaDetector() *JavaDetector {
	return &JavaDetector{
		baseDetector: baseDetector{
			name:     "Java",
			category: CategoryLanguage,
		},
	}
}

func (d *JavaDetector) Detect() Tool {
	result := command.Execute("java", "-version")
	tool := Tool{
		Name:     d.name,
		Category: string(d.Category()),
		Symbol:   "☕",
	}

	if !result.Exists {
		tool.Detected = false
		tool.Version = "Not installed"
		return tool
	}

	if result.Error != nil {
		tool.Detected = false
		if result.Output == "" {
			tool.Version = "Not installed"
		} else {
			tool.Version = parseVersion(result.Output)
		}
		return tool
	}

	path, _ := command.LookPath("java")
	tool.Path = path
	tool.Detected = true
	tool.Version = parseVersion(result.Output)
	return tool
}

func GetCoreDetectors() []Detector {
	return []Detector{
		NewGoDetector(),
		NewNodeDetector(),
		NewPythonDetector(),
		NewJavaDetector(),
	}
}
