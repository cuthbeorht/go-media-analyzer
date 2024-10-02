setup:
	@echo "Fetching dependencies"
	@go mod tidy

test-all:	
	@echo "Executing test suite"
	@go test -v ./...

publish:
	@echo "Publishing new version"
	sh ./build/cicd/scripts/publish.sh