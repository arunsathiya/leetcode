.PHONY: build install new test bench test-all clean watch help

# Extract positional arguments from MAKECMDGOALS
# Usage: make test 1          -> PROBLEM=1
#        make new 1 "Title" "Sig"  -> PROBLEM=1 TITLE="Title" SIGNATURE="Sig"
PROBLEM ?= $(word 2, $(MAKECMDGOALS))
TITLE ?= $(word 3, $(MAKECMDGOALS))
SIGNATURE ?= $(word 4, $(MAKECMDGOALS))

# Prevent Make from treating extra args as targets
%:
	@:

# Build the CLI tool
build:
	go build -trimpath -o leet ./cmd/leet

# Install the CLI tool to GOPATH/bin
install:
	go install ./cmd/leet

# Create a new problem
new:
	@if [ -z "$(PROBLEM)" ]; then echo "❌ Error: PROBLEM required. Usage: make new 123 \"Two Sum\" \"func twoSum(nums []int, target int) []int\""; exit 1; fi
	@if [ -z "$(TITLE)" ]; then echo "❌ Error: TITLE required. Usage: make new 123 \"Two Sum\" \"func twoSum(nums []int, target int) []int\""; exit 1; fi
	@if [ -z "$(SIGNATURE)" ]; then echo "❌ Error: SIGNATURE required. Usage: make new 123 \"Two Sum\" \"func twoSum(nums []int, target int) []int\""; exit 1; fi
	./leet new $(PROBLEM) "$(TITLE)" "$(SIGNATURE)"

# Run tests for a specific problem (optional)
test:
	@if [ -z "$(PROBLEM)" ]; then go test ./...; else ./leet test $(PROBLEM); fi

# Run benchmarks for a specific problem (optional)
bench:
	@if [ -z "$(PROBLEM)" ]; then go test -bench=. ./...; else ./leet bench $(PROBLEM); fi

# Run all tests
test-all:
	go test ./...

# Watch and test a specific problem on file changes
watch:
	@if [ -z "$(PROBLEM)" ]; then echo "❌ Error: PROBLEM required. Usage: make watch 1"; exit 1; fi
	@echo "👀 Watching problem $(PROBLEM)..."
	@while inotifywait -e modify src/$(PROBLEM)-*/main.go src/$(PROBLEM)-*/main_test.go 2>/dev/null; do clear && make test $(PROBLEM); done

# Clean build artifacts
clean:
	rm -f leet

# Help
help:
	@echo "LeetCode Go Makefile"
	@echo ""
	@echo "Commands:"
	@echo "  make build                      Build the CLI tool"
	@echo "  make install                    Install CLI tool to GOPATH/bin"
	@echo "  make new N TITLE SIGNATURE      Create new problem"
	@echo "  make test [N]                   Run tests (specific problem or all)"
	@echo "  make bench [N]                  Run benchmarks (specific problem or all)"
	@echo "  make watch N                    Watch & auto-test on file changes"
	@echo "  make test-all                   Run all tests"
	@echo "  make clean                      Clean build artifacts"
	@echo ""
	@echo "Examples:"
	@echo "  make build"
	@echo "  make new 1 \"Two Sum\" \"func twoSum(nums []int, target int) []int\""
	@echo "  make test 1"
	@echo "  make test                       (run all tests)"
	@echo "  make watch 1"
	@echo "  make bench 1"
	@echo "  make test-all"