.PHONY: build install new test bench clean test-all help

LANGUAGE ?= go

help:
	@echo "LeetCode Multi-Language Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make build [LANGUAGE=go]                     Build CLI tool (Go only)"
	@echo "  make install [LANGUAGE=go]                   Install CLI tool"
	@echo "  make new PROBLEM=N TITLE=\"...\" SIGNATURE=\"...\" [LANGUAGE=go]   Create new problem"
	@echo "  make test [PROBLEM=N] [LANGUAGE=go]          Run tests"
	@echo "  make bench [PROBLEM=N] [LANGUAGE=go]         Run benchmarks (Go only)"
	@echo "  make clean [LANGUAGE=go]                     Clean build artifacts"
	@echo "  make test-all [LANGUAGE=go]                  Run all tests in language"
	@echo ""
	@echo "Examples:"
	@echo "  make new PROBLEM=1 TITLE=\"Two Sum\" SIGNATURE=\"func twoSum(nums []int, target int) []int\""
	@echo "  make new PROBLEM=1 TITLE=\"Two Sum\" LANGUAGE=go"
	@echo "  make test PROBLEM=1"
	@echo "  make test-all"

build:
	cd $(LANGUAGE) && make build

install:
	cd $(LANGUAGE) && make install

new:
	cd $(LANGUAGE) && make new PROBLEM=$(PROBLEM) TITLE="$(TITLE)" SIGNATURE="$(SIGNATURE)"

test:
ifdef PROBLEM
	cd $(LANGUAGE) && make test PROBLEM=$(PROBLEM)
else
	cd $(LANGUAGE) && make test
endif

bench:
ifdef PROBLEM
	cd $(LANGUAGE) && make bench PROBLEM=$(PROBLEM)
else
	cd $(LANGUAGE) && make bench
endif

clean:
	cd $(LANGUAGE) && make clean

test-all:
	cd $(LANGUAGE) && make test-all
