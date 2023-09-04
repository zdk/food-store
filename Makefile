GOBIN ?= $$(go env GOPATH)/bin

.PHONY: check-coverage
check-coverage:
	go test -race -coverprofile=coverage.txt -covermode=atomic
