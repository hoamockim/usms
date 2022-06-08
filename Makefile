GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

update:
	go mod tidy
all: build
build:
	$(GOBUILD) -v -ldflags="-extldflags=-static" -o "usms" cmd/profile/main.go

test:
	$(GOTEST) -v $(go list ./.. | grep -v vendor) -covermode=count -coverprofile=sample.cov

run internal:
	go run app/cmd/internal/*.go
