# Test LeetCode Problem

Run tests for a specific LeetCode problem solution.

## Usage
```
/lc:test <problem-directory>
```

Example: `/lc:test 1-two-sum`

## Description
This command runs tests for a specific problem using Go's built-in testing framework.

## Implementation
```bash
cd "src/$ARGUMENTS" && go test -v
```

This will run all tests in the specified problem directory and show detailed output.