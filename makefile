new:
	@echo "Creating new problem folder: $(NAME)"
	@mkdir -p $(shell echo $(NAME) | tr -d '.' | tr 'A-Z ' 'a-z-' )
	@echo "package main\n\nfunc main() {}\n" > $(shell echo $(NAME) | tr -d '.' | tr 'A-Z ' 'a-z-' )/main.go
	@echo "package main\n\nimport \"testing\"\n\nfunc TestMain(t *testing.T) {}\n" > $(shell echo $(NAME) | tr -d '.' | tr 'A-Z ' 'a-z-' )/main_test.go

title:
	@branch=$$(git rev-parse --abbrev-ref HEAD); \
	title=$$(echo $$branch | sed -E 's/solve\/([0-9]+)-/\1. /' | sed 's/-/ /g' | awk 'BEGIN{FS=OFS=" "}{for(i=2;i<=NF;i++)$$i=toupper(substr($$i,1,1))tolower(substr($$i,2))}1'); \
	echo "new(problem): $$title" | pbcopy