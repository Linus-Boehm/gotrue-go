
test: ## Run tests.
	go test $(CHECK_FILES) -coverprofile=coverage.out -coverpkg ./... -p 1 -race -v -count=1