build-run:
	go build -o ./tmp/main ./cmd/app/main.go
	./tmp/main

watch:
	@if command -v wgo > /dev/null; then \
			echo "Watching..."; \
			wgo -file .go clear :: wgo run ./cmd/app/main.go; \
		else \
			read -p "'wgo' is not installed. Do you want to install it? [Y/n] " choice; \
			if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
				go install github.com/bokwoon95/wgo@latest; \
				wgo -file .go clear :: wgo run ./cmd/app/main.go; \
			else \
				echo "You chose not to install wgo. Exiting..."; \
				exit 1; \
			fi; \
    	fi

test:
	go test ./...

lint:
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run ./...; \
	else \
		read -p "'golangci-lint' is not installed. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
			golangci-lint run ./...; \
		else \
			echo "You chose not to install golangci-lint. Exiting..."; \
			exit 1; \
		fi; \
	fi

.PHONY: watch build-run test lint