package command

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var DefaultTimeout = 5 * time.Second

type Result struct {
	Command string
	Output  string
	Error   error
	Exists  bool
}

func Execute(name string, args ...string) Result {
	return ExecuteWithTimeout(name, DefaultTimeout, args...)
}

func ExecuteWithTimeout(name string, timeout time.Duration, args ...string) Result {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	output, err := cmd.CombinedOutput()

	exists := true
	if err != nil {
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			exists = false
		}
	}

	return Result{
		Command: name,
		Output:  strings.TrimSpace(string(output)),
		Error:   err,
		Exists:  exists,
	}
}

func LookPath(name string) (string, error) {
	path, err := exec.LookPath(name)
	if err != nil {
		return "", fmt.Errorf("%s not found in PATH", name)
	}
	return path, nil
}

func Getenv(key string) string {
	return os.Getenv(key)
}

func GetenvOrDefault(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}
