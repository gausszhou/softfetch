package detect

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Tool struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Path     string `json:"path,omitempty"`
	Detected bool   `json:"detected"`
	Symbol   string `json:"symbol"`
	Category string `json:"category"`
}

type DetectionResult struct {
	Tools    []Tool    `json:"tools"`
	OS       string    `json:"os"`
	Arch     string    `json:"arch"`
	Detected time.Time `json:"detected_at"`
}

type Category string

const (
	CategoryLanguage   Category = "Language"
	CategoryPackageMgr Category = "Package Manager"
	CategoryBuildTool  Category = "Build Tool"
	CategoryCompiler   Category = "Compiler"
	CategoryRuntime    Category = "Runtime"
	CategoryOther      Category = "Other"
)

func (t *Tool) String() string {
	if t.Detected {
		return fmt.Sprintf("%s (%s)", t.Version, t.Path)
	}
	return "Not detected"
}

type Detector interface {
	Detect() Tool
	Name() string
	Category() Category
}

type baseDetector struct {
	name     string
	category Category
}

func (b *baseDetector) Name() string       { return b.name }
func (b *baseDetector) Category() Category { return b.category }

func Detect(detectors ...Detector) DetectionResult {
	result := DetectionResult{
		Detected: time.Now(),
		OS:       runtime.GOOS,
		Arch:     runtime.GOARCH,
	}

	if len(detectors) == 0 {
		return result
	}

	resultChan := make(chan Tool, len(detectors))
	var wg sync.WaitGroup

	for _, d := range detectors {
		wg.Add(1)
		go func(detector Detector) {
			defer wg.Done()
			resultChan <- detector.Detect()
		}(d)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for tool := range resultChan {
		result.Tools = append(result.Tools, tool)
	}

	return result
}

func parseVersion(output string) string {
	output = strings.TrimSpace(output)
	output = strings.Trim(output, "\n")

	lines := strings.Split(output, "\n")
	if len(lines) > 0 {
		firstLine := lines[0]
		firstLine = strings.ReplaceAll(firstLine, "\r", "")

		re := regexp.MustCompile(`go\s*(\d+\.\d+[\w.-]*)`)
		matches := re.FindStringSubmatch(firstLine)
		if len(matches) > 1 {
			return matches[1]
		}

		re = regexp.MustCompile(`Python\s+(\d+\.\d+\.\d+)`)
		matches = re.FindStringSubmatch(firstLine)
		if len(matches) > 1 {
			return matches[1]
		}

		re = regexp.MustCompile(`openjdk\s+version\s+"?(\d+[^"]+)"?`)
		matches = re.FindStringSubmatch(firstLine)
		if len(matches) > 1 {
			return matches[1]
		}

		if idx := strings.Index(firstLine, "version"); idx != -1 {
			rest := firstLine[idx+len("version"):]
			rest = strings.TrimSpace(rest)
			rest = strings.Trim(rest, "\"")

			for _, sep := range []string{" ", ":", "v"} {
				if strings.HasPrefix(rest, sep) {
					rest = strings.TrimPrefix(rest, sep)
				}
			}

			fields := strings.Fields(rest)
			if len(fields) > 0 {
				return fields[0]
			}
		}

		re = regexp.MustCompile(`(\d+\.\d+\.\d+[\w.-]*)`)
		matches = re.FindStringSubmatch(firstLine)
		if len(matches) > 1 {
			return matches[1]
		}

		return firstLine
	}

	return output
}
