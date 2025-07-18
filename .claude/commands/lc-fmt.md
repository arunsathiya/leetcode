# Format Go Code

Format all Go code in the project using gofmt.

## Usage
```
/lc:fmt
```

## Description
This command formats all Go code in the project according to the standard Go formatting rules.

## Implementation
```bash
gofmt -w src/*/main.go src/*/main_test.go
```

This ensures consistent code style across all problem solutions.