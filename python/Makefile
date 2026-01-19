.PHONY: install test bench clean help

help:
	@echo "Python Solutions Makefile"
	@echo ""
	@echo "Commands:"
	@echo "  make install                    Install dependencies"
	@echo "  make test [PROBLEM=N]           Run tests"
	@echo "  make test-all                   Run all tests"
	@echo "  make clean                      Clean build artifacts"

install:
	pip install -e .

test:
ifdef PROBLEM
	pytest solutions/test_$(PROBLEM).py -v
else
	pytest -v
endif

test-all:
	pytest -v

clean:
	rm -rf build dist *.egg-info __pycache__ .pytest_cache
