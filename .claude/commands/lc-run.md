# Run LeetCode Problem

Run a specific LeetCode problem solution.

## Usage
```
/lc:run <problem-directory>
```

Example: `/lc:run 1-two-sum`

## Description
This command compiles and runs a specific problem's main function, useful for manual testing or debugging.

## Implementation
```bash
cd "src/$ARGUMENTS" && go run main.go
```

This will compile and execute the main function in the specified problem directory.