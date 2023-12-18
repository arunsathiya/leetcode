new:
	@echo "Creating project: $(NAME)"
	@mkdir -p $(shell echo $(NAME) | tr -d '.' | tr 'A-Z ' 'a-z-' )
	@echo "package main\n\nfunc main() {}\n" > $(shell echo $(NAME) | tr -d '.' | tr 'A-Z ' 'a-z-' )/main.go
	@echo "package main\n\nimport \"testing\"\n\nfunc TestMain(t *testing.T) {}\n" > $(shell echo $(NAME) | tr -d '.' | tr 'A-Z ' 'a-z-' )/main_test.go