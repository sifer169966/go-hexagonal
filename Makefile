GO_BINARY_NAME=go-hex # <- change this value to your binary name
VERSION=$(shell git describe --tags || git rev-parse --short HEAD || echo "unknown version")
LDFLAGS+= -X "hexagonal-template/cmd/cmds.Version=$(VERSION)"
LDFLAGS+= -X "hexagonal-template/cmd/cmds.GoVersion=$(shell go version | sed -r 's/go version go(.*)\ .*/\1/')"



# Always turn on go module when use `go` command.
GO := GO111MODULE=on go

# GO build preparation
.PHONY: ci
ci:
	$(GO) mod download && \
	$(GO) mod verify && \
	$(GO) mod vendor && \
	$(GO) fmt ./... \

# Build GO application
# -mod=vendor 
# tells the go command to use the vendor directory. In this mode,
# the go command will not use the network or the module cache.
# -v
# print the names of packages as they are compiled.
# -a
# force rebuilding of packages that are already up-to-date.
# -o
# -ldsflags
# tells the version and go version.
.PHONY: build
build:
	$(GO) build -mod=vendor -ldflags '$(LDFLAGS)' -a -v -o $(GO_BINARY_NAME) ./cmd/main.go

start:
	go run ./cmd/main.go serve-rest

.PHONY: test
test:
	go test ./... -coverprofile coverage.out

# Clean up when build the application on local directory.
.PHONY: clean
clean:
	@rm -rf $(GO_BINARY_NAME) ./vendor