package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

func NewProblem(number int, title, signature string, createBranch, useWorktree bool) error {
	// Convert title to kebab-case
	dirName := fmt.Sprintf("%d-%s", number, toKebabCase(title))
	problemDir := filepath.Join("src", dirName)
	
	// Check if directory already exists
	if _, err := os.Stat(problemDir); !os.IsNotExist(err) {
		return fmt.Errorf("problem directory '%s' already exists", problemDir)
	}
	
	// Create directory
	if err := os.MkdirAll(problemDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	
	// Extract function name from signature
	funcName := extractFunctionName(signature)
	if funcName == "" {
		return fmt.Errorf("could not extract function name from signature: %s", signature)
	}
	
	// Generate main.go
	mainContent := fmt.Sprintf(`package main

%s {
	// TODO: implement solution
	panic("not implemented")
}
`, signature)
	
	mainPath := filepath.Join(problemDir, "main.go")
	if err := os.WriteFile(mainPath, []byte(mainContent), 0644); err != nil {
		return fmt.Errorf("failed to write main.go: %v", err)
	}
	
	// Generate test file
	testData := struct {
		Number   int
		Title    string
		FuncName string
	}{
		Number:   number,
		Title:    title,
		FuncName: funcName,
	}
	
	testTemplate := `package main

import (
	"testing"
)

func Test{{.FuncName}}(t *testing.T) {
	tests := []struct {
		name string
		// TODO: add test fields
		want interface{}
	}{
		{
			name: "Example 1",
			// TODO: add test case
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: implement test
			// got := {{.FuncName}}(...)
			// if got != tt.want {
			//     t.Errorf("{{.FuncName}}() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Benchmark{{.FuncName}}(b *testing.B) {
	// TODO: add benchmark inputs
	for i := 0; i < b.N; i++ {
		// {{.FuncName}}(...)
	}
}
`
	
	tmpl, err := template.New("test").Parse(testTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse test template: %v", err)
	}
	
	testPath := filepath.Join(problemDir, "main_test.go")
	testFile, err := os.Create(testPath)
	if err != nil {
		return fmt.Errorf("failed to create test file: %v", err)
	}
	defer testFile.Close()
	
	if err := tmpl.Execute(testFile, testData); err != nil {
		return fmt.Errorf("failed to execute test template: %v", err)
	}
	
	fmt.Printf("✅ Created problem %d: %s\n", number, title)
	fmt.Printf("   Directory: %s\n", problemDir)
	fmt.Printf("   Files: main.go, main_test.go\n")
	
	// Create branch if requested
	if createBranch {
		// Use the same kebab-case function for consistency
		kebabTitle := toKebabCase(title)
		branchName := fmt.Sprintf("solve/%d-%s", number, kebabTitle)
		
		if err := CreateBranch(branchName, useWorktree); err != nil {
			fmt.Printf("⚠️  Warning: Failed to create branch: %v\n", err)
			fmt.Printf("   You can create it manually with: leet branch %d \"%s\"\n", number, title)
		}
	}
	
	return nil
}

func toKebabCase(s string) string {
	// Convert to lowercase and replace spaces/special chars with hyphens
	s = strings.ToLower(s)
	
	// Replace multiple spaces/special chars with single hyphen
	re := regexp.MustCompile(`[^a-z0-9]+`)
	s = re.ReplaceAllString(s, "-")
	
	// Remove leading/trailing hyphens
	s = strings.Trim(s, "-")
	
	return s
}

func extractFunctionName(signature string) string {
	// Extract function name from signature like "func twoSum(...) ..."
	re := regexp.MustCompile(`func\s+(\w+)`)
	matches := re.FindStringSubmatch(signature)
	if len(matches) >= 2 {
		return matches[1]
	}
	return ""
}