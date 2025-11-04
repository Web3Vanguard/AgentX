run:
	@go build -o bin/app ./cmd
	@./bin/app
linux-build:
	@GOOS=linux GOARCH=amd64 go build -o bin/agentX-linux-amd64 ./cmd/main.go
windows-build:
	@GOOS=windows GOARCH=amd64 go build -o bin/agentX-windows-amd64.exe ./cmd/main.go
mac-build:
	@GOOS=darwin GOARCH=amd64 go build -o bin/agentX-darwin-amd64 ./cmd/main.go
.PHONY: run
