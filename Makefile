setup:
	@echo "Fetching dependencies"
	@go mod tidy

build:
	@echo "Building binary"
	@go build -o media-analyzer main.go

test:	
	@echo "Executing test suite"
	@go test ./...

run:
	@echo "Running media analyzer"

