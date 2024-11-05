# .PHONY: lint
# lint:
# 	# Run golangci-lint on each Go module in the repository
# 	cd app1 && golangci-lint run --fix
# 	cd app2 && golangci-lint run --fix

.PHONY: ci-checks
ci-checks:
	# Format, vet, tidy, build, and test each Go module in the repository
	cd app1 && gofmt -s -w . && go vet ./... && go mod tidy && go build -o /dev/null ./... && go test -v ./...
	cd app2 && gofmt -s -w . && go vet ./... && go mod tidy && go build -o /dev/null ./... && go test -v ./...
