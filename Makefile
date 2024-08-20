setup:
	@echo "Fetching dependencies"
	@go mod tidy

build:
	@echo "Building binary"
	@go build -o media-analyzer main.go

test:
	@echo "Setting up tests"
	cp -R fixtures /tmp
	@echo "Executing test suite"
	@go test ./...
	@echo "Cleaning up after myself"
	rm -rf /tmp/fixtures

run:
	@echo "Running media analyzer"

