# Create New LeetCode Problem

Create a new LeetCode problem solution with proper Go structure.

## Usage
```
/lc:new <problem-number> <problem-name>
```

Example: `/lc:new 567 permutation-in-string`

## Description
This command creates a new LeetCode problem solution by:
1. Creating a new directory in `src/` with the format `{number}-{problem-name}`
2. Creating `main.go` with a basic function template
3. Creating `main_test.go` with a test template
4. Opening the generated files for editing

## Implementation
```bash
mkdir -p "src/$1-$2"
cd "src/$1-$2"

# Create main.go with basic template
cat > main.go << 'EOF'
package main

// TODO: Implement the solution function here
func solution() {
	// Your implementation here
}

func main() {
	// Optional: Add a main function for manual testing
}
EOF

# Create main_test.go with test template
cat > main_test.go << 'EOF'
package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		// TODO: Add test case fields
		want interface{}
	}{
		{
			name: "Example 1",
			// TODO: Add test case values
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: Call your solution function and compare with tt.want
			// if got := solution(); got != tt.want {
			//     t.Errorf("solution() = %v, want %v", got, tt.want)
			// }
		})
	}
}
EOF

echo "Created new problem: src/$1-$2"
echo "Files created:"
echo "  - src/$1-$2/main.go"
echo "  - src/$1-$2/main_test.go"
```

After running this command:
1. Implement your solution in the generated `main.go` file
2. Add appropriate test cases in `main_test.go`
3. Use `/lc:test` to run your tests