run:
	@go build -o bin/app ./cmd
	@./bin/app

.PHONY: run