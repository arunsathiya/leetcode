.PHONY: install test bench clean help

help:
	@echo "TypeScript Solutions Makefile"
	@echo ""
	@echo "Commands:"
	@echo "  make install                    Install dependencies"
	@echo "  make test [PROBLEM=N]           Run tests"
	@echo "  make test-all                   Run all tests"
	@echo "  make clean                      Clean build artifacts"

install:
	npm install

test:
ifdef PROBLEM
	npm test -- $(PROBLEM)
else
	npm test
endif

test-all:
	npm test

clean:
	rm -rf node_modules dist
