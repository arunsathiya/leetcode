package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func RunTests(problem string) error {
	if problem == "" {
		// Run all tests
		fmt.Println("🧪 Running all tests...")
		cmd := exec.Command("go", "test", "./...")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
	
	// Run specific problem test
	problemDir, err := findProblemDir(problem)
	if err != nil {
		return err
	}
	
	fmt.Printf("🧪 Running tests for problem %s...\n", problem)
	cmd := exec.Command("go", "test", "-v")
	cmd.Dir = problemDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func RunBenchmarks(problem string) error {
	if problem == "" {
		// Run all benchmarks
		fmt.Println("⚡ Running all benchmarks...")
		cmd := exec.Command("go", "test", "-bench=.", "./...")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
	
	// Run specific problem benchmark
	problemDir, err := findProblemDir(problem)
	if err != nil {
		return err
	}
	
	fmt.Printf("⚡ Running benchmarks for problem %s...\n", problem)
	cmd := exec.Command("go", "test", "-bench=.", "-benchmem")
	cmd.Dir = problemDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func findProblemDir(problem string) (string, error) {
	// Convert problem to int to handle both "123" and numeric input
	problemNum, err := strconv.Atoi(problem)
	if err != nil {
		return "", fmt.Errorf("invalid problem number: %s", problem)
	}
	
	// Find directory that starts with the problem number
	srcDir := "src"
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return "", fmt.Errorf("could not read src directory: %v", err)
	}
	
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		
		// Check if directory name starts with the problem number
		dirName := entry.Name()
		if strings.HasPrefix(dirName, fmt.Sprintf("%d-", problemNum)) {
			return filepath.Join(srcDir, dirName), nil
		}
	}
	
	return "", fmt.Errorf("problem %d not found", problemNum)
}