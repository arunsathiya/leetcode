# Test All LeetCode Problems

Run tests for all LeetCode problems in the project.

## Usage
```
/lc:test-all
```

## Description
This command runs tests for all problems in the src/ directory using Go's testing framework.

## Implementation
```bash
cd src && for dir in */; do echo "Testing $dir..." && cd "$dir" && go test -v && cd ..; done
```

This will iterate through all problem directories and run their tests, showing results for each.