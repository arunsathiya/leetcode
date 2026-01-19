.PHONY: build install new test bench test-all clean

# Build the CLI tool
build:
	go build -trimpath -o leet ./cmd/leet

# Install the CLI tool to GOPATH/bin
install:
	go install ./cmd/leet

# Create a new problem (requires PROBLEM, TITLE, and SIGNATURE variables)
new:
ifndef PROBLEM
	$(error PROBLEM is required. Usage: make new PROBLEM=123 TITLE="Two Sum" SIGNATURE="func twoSum(nums []int, target int) []int")
endif
ifndef TITLE
	$(error TITLE is required. Usage: make new PROBLEM=123 TITLE="Two Sum" SIGNATURE="func twoSum(nums []int, target int) []int")
endif
ifndef SIGNATURE
	$(error SIGNATURE is required. Usage: make new PROBLEM=123 TITLE="Two Sum" SIGNATURE="func twoSum(nums []int, target int) []int")
endif
	./leet new $(PROBLEM) "$(TITLE)" "$(SIGNATURE)"

# Run tests for a specific problem (requires PROBLEM variable)
test:
ifdef PROBLEM
	./leet test $(PROBLEM)
else
	./leet test
endif

# Run benchmarks for a specific problem (requires PROBLEM variable)
bench:
ifdef PROBLEM
	./leet bench $(PROBLEM)
else
	./leet bench
endif

# Run all tests
test-all:
	go test ./...

# Clean build artifacts
clean:
	rm -f leet

# Help
help:
	@echo "LeetCode Project Makefile"
	@echo ""
	@echo "Commands:"
	@echo "  make build                                    Build the CLI tool"
	@echo "  make install                                  Install CLI tool to GOPATH/bin"
	@echo "  make new PROBLEM=123 TITLE=\"Two Sum\" SIGNATURE=\"func twoSum(nums []int, target int) []int\""
	@echo "                                                Create new problem"
	@echo "  make test [PROBLEM=123]                       Run tests (specific or all)"
	@echo "  make bench [PROBLEM=123]                      Run benchmarks (specific or all)"
	@echo "  make test-all                                 Run all tests"
	@echo "  make clean                                    Clean build artifacts"
	@echo ""
	@echo "Examples:"
	@echo "  make build"
	@echo "  make new PROBLEM=1 TITLE=\"Two Sum\" SIGNATURE=\"func twoSum(nums []int, target int) []int\""
	@echo "  make test PROBLEM=1"
	@echo "  make bench PROBLEM=1"
	@echo "  make test-all"