package main

import (
	"fmt"
	"os"

	"github.com/francisco3ferraz/go-nuts/internal/metrics"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: metrics-analyzer <input-file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	analyzer := metrics.NewAnalyzer()

	summary, err := analyzer.AnalyzeFile(inputFile)
	if err != nil {
		fmt.Printf("analysis failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(summary)
}
