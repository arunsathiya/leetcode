[![Checks](https://github.com/arunsathiya/leetcode/actions/workflows/tests.yml/badge.svg)](https://github.com/arunsathiya/leetcode/actions)

# LeetCode Solutions - Multi-Language

This repository contains solutions for LeetCode problems in multiple programming languages:
- **Go** - Primary focus
- **TypeScript** - Coming soon
- **Python** - Coming soon

## Project Structure

```
leetcode/
├── go/              # Go solutions
├── typescript/      # TypeScript solutions
├── python/          # Python solutions
└── Makefile         # Root makefile for multi-language support
```

## Quick Start (Go)

```bash
# Build the CLI tool
make build LANG=go

# Create a new problem
make new PROBLEM=1 TITLE="Two Sum" SIGNATURE="func twoSum(nums []int, target int) []int" LANG=go

# Run tests
make test LANG=go
make test PROBLEM=1 LANG=go

# Run benchmarks
make bench LANG=go
```

## Language-Specific

### Go
Navigate to the `go/` directory for Go-specific commands:

```bash
cd go/
make build
make new PROBLEM=1 TITLE="Two Sum" SIGNATURE="func twoSum(nums []int, target int) []int"
make test
make bench
```