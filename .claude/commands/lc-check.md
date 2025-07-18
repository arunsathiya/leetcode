# Check Code Quality

Run code quality checks for all Go code in the project.

## Usage
```
/lc:check
```

## Description
This command runs Go's built-in tools to check code quality, including vet for potential issues and formatting checks.

## Implementation
```bash
gofmt -d src/*/main.go src/*/main_test.go && go vet ./src/...
```

This helps catch potential issues and ensures code follows Go best practices.