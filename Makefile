BRANCH=`git rev-parse --abbrev-ref HEAD`
COMMIT=`git rev-parse --short HEAD`
VERSION=`git describe --always --tags --dirty=-hacky`
GOLDFLAGS="-X main.branch=$(BRANCH) -X main.commit=$(COMMIT) -X main.version=$(VERSION)"
PACKAGENAME=`go list .`
all: test build
setup:
	@echo "-> install build deps"
	@go get -u "golang.org/x/tools/cmd/vet"
vet:
	@echo "-> go vet"
	@go vet ./...
fmt:
	@echo "-> go fmt"
	@go fmt ./...
install: test
	@echo "-> go install"
	@go install -ldflags=$(GOLDFLAGS)
build:
	@echo "-> go build"
	@go build -ldflags=$(GOLDFLAGS)
linux: test
	GOARCH=amd64 GOOS=linux go build -ldflags=$(GOLDFLAGS)
test: fmt vet errcheck
	@echo "-> go test"
	@go test ./... -cover -v
.PHONY: setup errcheck vet fmt install build test
