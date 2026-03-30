package detect

import (
	"github.com/gausszhou/softfetch/internal/command"
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
	tool := Tool{
		Name:     d.name,
		Category: string(d.Category()),
		Symbol:   "⬢",
	}

	path, err := command.LookPath("go")
	if err != nil {
		tool.Detected = false
		tool.Version = "Not installed"
		return tool
	}

	result := command.Execute("go", "version")
	if result.Error != nil {
		tool.Detected = false
		tool.Version = parseVersion(result.Output)
		return tool
	}

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
	tool := Tool{
		Name:     d.name,
		Category: string(d.Category()),
		Symbol:   "⬢",
	}

	path, err := command.LookPath("node")
	if err != nil {
		tool.Detected = false
		tool.Version = "Not installed"
		return tool
	}

	result := command.Execute("node", "--version")
	if result.Error != nil {
		tool.Detected = false
		tool.Version = parseVersion(result.Output)
		return tool
	}

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
			path, err := command.LookPath(cmd)
			if err == nil {
				tool.Path = path
			}
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
	tool := Tool{
		Name:     d.name,
		Category: string(d.Category()),
		Symbol:   "☕",
	}

	path, err := command.LookPath("java")
	if err != nil {
		tool.Detected = false
		tool.Version = "Not installed"
		return tool
	}

	result := command.Execute("java", "-version")
	if result.Error != nil {
		tool.Detected = false
		if result.Output == "" {
			tool.Version = "Not installed"
		} else {
			tool.Version = parseVersion(result.Output)
		}
		return tool
	}

	tool.Path = path
	tool.Detected = true
	tool.Version = parseVersion(result.Output)
	return tool
}

type CDetector struct {
	baseDetector
}

func NewCDetector() *CDetector {
	return &CDetector{
		baseDetector: baseDetector{
			name:     "C",
			category: CategoryLanguage,
		},
	}
}

func (d *CDetector) Detect() Tool {
	tool := Tool{
		Name:     d.name,
		Category: string(d.Category()),
		Symbol:   "🔧",
	}

	commands := []string{"gcc", "clang", "cc"}
	for _, cmd := range commands {
		path, err := command.LookPath(cmd)
		if err != nil {
			continue
		}
		result := command.Execute(cmd, "--version")
		if result.Error == nil {
			tool.Detected = true
			tool.Version = parseVersion(result.Output)
			tool.Path = path
			return tool
		}
	}

	tool.Detected = false
	tool.Version = "Not installed"
	return tool
}

type CppDetector struct {
	baseDetector
}

func NewCppDetector() *CppDetector {
	return &CppDetector{
		baseDetector: baseDetector{
			name:     "C++",
			category: CategoryLanguage,
		},
	}
}

func (d *CppDetector) Detect() Tool {
	tool := Tool{
		Name:     d.name,
		Category: string(d.Category()),
		Symbol:   "🔧",
	}

	commands := []string{"g++", "clang++", "c++"}
	for _, cmd := range commands {
		path, err := command.LookPath(cmd)
		if err != nil {
			continue
		}
		result := command.Execute(cmd, "--version")
		if result.Error == nil {
			tool.Detected = true
			tool.Version = parseVersion(result.Output)
			tool.Path = path
			return tool
		}
	}

	tool.Detected = false
	tool.Version = "Not installed"
	return tool
}

type RustDetector struct {
	baseDetector
}

func NewRustDetector() *RustDetector {
	return &RustDetector{
		baseDetector: baseDetector{
			name:     "Rust",
			category: CategoryLanguage,
		},
	}
}

func (d *RustDetector) Detect() Tool {
	tool := Tool{
		Name:     d.name,
		Category: string(d.Category()),
		Symbol:   "🦀",
	}

	path, err := command.LookPath("rustc")
	if err != nil {
		tool.Detected = false
		tool.Version = "Not installed"
		return tool
	}

	result := command.Execute("rustc", "--version")
	if result.Error == nil {
		tool.Detected = true
		tool.Version = parseVersion(result.Output)
		tool.Path = path
		return tool
	}

	tool.Detected = false
	tool.Version = "Not installed"
	return tool
}

type PHPDetector struct {
	baseDetector
}

func NewPHPDetector() *PHPDetector {
	return &PHPDetector{
		baseDetector: baseDetector{
			name:     "PHP",
			category: CategoryLanguage,
		},
	}
}

func (d *PHPDetector) Detect() Tool {
	tool := Tool{
		Name:     d.name,
		Category: string(d.Category()),
		Symbol:   "🐘",
	}

	path, err := command.LookPath("php")
	if err != nil {
		tool.Detected = false
		tool.Version = "Not installed"
		return tool
	}

	result := command.Execute("php", "--version")
	if result.Error != nil {
		tool.Detected = false
		tool.Version = parseVersion(result.Output)
		return tool
	}

	tool.Path = path
	tool.Detected = true
	tool.Version = parseVersion(result.Output)
	return tool
}

type DockerDetector struct {
	baseDetector
}

func NewDockerDetector() *DockerDetector {
	return &DockerDetector{
		baseDetector: baseDetector{
			name:     "Docker",
			category: CategoryRuntime,
		},
	}
}

func (d *DockerDetector) Detect() Tool {
	tool := Tool{
		Name:     d.name,
		Category: string(d.Category()),
		Symbol:   "🐳",
	}

	path, err := command.LookPath("docker")
	if err != nil {
		tool.Detected = false
		tool.Version = "Not installed"
		return tool
	}

	result := command.Execute("docker", "--version")
	if result.Error != nil {
		tool.Detected = false
		tool.Version = parseVersion(result.Output)
		return tool
	}

	tool.Path = path
	tool.Detected = true
	tool.Version = parseVersion(result.Output)
	return tool
}

type GitDetector struct {
	baseDetector
}

func NewGitDetector() *GitDetector {
	return &GitDetector{
		baseDetector: baseDetector{
			name:     "Git",
			category: CategoryBuildTool,
		},
	}
}

func (d *GitDetector) Detect() Tool {
	tool := Tool{
		Name:     d.name,
		Category: string(d.Category()),
		Symbol:   "📂",
	}

	path, err := command.LookPath("git")
	if err != nil {
		tool.Detected = false
		tool.Version = "Not installed"
		return tool
	}

	result := command.Execute("git", "--version")
	if result.Error != nil {
		tool.Detected = false
		tool.Version = parseVersion(result.Output)
		return tool
	}

	tool.Path = path
	tool.Detected = true
	tool.Version = parseVersion(result.Output)
	return tool
}

type RubyDetector struct {
	baseDetector
}

func NewRubyDetector() *RubyDetector {
	return &RubyDetector{
		baseDetector: baseDetector{
			name:     "Ruby",
			category: CategoryLanguage,
		},
	}
}

func (d *RubyDetector) Detect() Tool {
	tool := Tool{
		Name:     d.name,
		Category: string(d.Category()),
		Symbol:   "💎",
	}

	path, err := command.LookPath("ruby")
	if err != nil {
		tool.Detected = false
		tool.Version = "Not installed"
		return tool
	}

	result := command.Execute("ruby", "--version")
	if result.Error != nil {
		tool.Detected = false
		tool.Version = parseVersion(result.Output)
		return tool
	}

	tool.Path = path
	tool.Detected = true
	tool.Version = parseVersion(result.Output)
	return tool
}

type DotNetDetector struct {
	baseDetector
}

func NewDotNetDetector() *DotNetDetector {
	return &DotNetDetector{
		baseDetector: baseDetector{
			name:     ".NET",
			category: CategoryRuntime,
		},
	}
}

func (d *DotNetDetector) Detect() Tool {
	tool := Tool{
		Name:     d.name,
		Category: string(d.Category()),
		Symbol:   "🔷",
	}

	path, err := command.LookPath("dotnet")
	if err != nil {
		tool.Detected = false
		tool.Version = "Not installed"
		return tool
	}

	result := command.Execute("dotnet", "--version")
	if result.Error != nil {
		tool.Detected = false
		tool.Version = parseVersion(result.Output)
		return tool
	}

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
		NewCDetector(),
		NewCppDetector(),
		NewRustDetector(),
		NewPHPDetector(),
		NewDockerDetector(),
		NewGitDetector(),
		NewRubyDetector(),
		NewDotNetDetector(),
	}
}
