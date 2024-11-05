.PHONY: lint
lint:
	# Run golangci-lint on each Go module in the repository
	cd app1 && golangci-lint run --fix
	cd app2 && golangci-lint run --fix
