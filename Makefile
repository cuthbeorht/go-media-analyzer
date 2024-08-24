setup:
	@echo "Fetching dependencies"
	@go mod tidy

test:	
	@echo "Executing test suite"
	@go test ./...

run:
	@echo "Running media analyzer"

