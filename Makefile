setup:
	@echo "Fetching dependencies"
	@go mod tidy

test:	
	@echo "Executing test suite"
	@go test ./...

publish:
	@echo "Publishing new version"
	sh ./build/cicd/scripts/publish.sh