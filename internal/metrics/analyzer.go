package metrics

import (
	"fmt"
	"os"
	"strings"
)

type Analyzer struct{}

func NewAnalyzer() Analyzer {
	return Analyzer{}
}

func (a Analyzer) AnalyzeFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(content), "\n")
	nonEmpty := 0
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			nonEmpty++
		}
	}

	return fmt.Sprintf("file=%s total_lines=%d non_empty_lines=%d", path, len(lines), nonEmpty), nil
}
