VERSION := $(AppVersion)
COMMIT := $(shell git rev-parse --short HEAD)
BUILDDATE := $(shell date +%Y-%m-%d)

LDFLAGS := -X 'main.AppVersion=$(VERSION)' -X 'main.CommitHash=$(COMMIT)' -X 'main.BuildDate=$(BUILDDATE)'


install:
	go mod tidy

test:
	go test -v ./...  -race -coverprofile=coverage.out -covermode=atomic

run: build
	./chatz "This is test."

build:
	go build -ldflags="$(LDFLAGS)" -o chatz .

build-all:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o build/chatz-linux-amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o build/chatz-linux-arm64
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="$(LDFLAGS)" -o build/chatz-linux-arm
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o build/chatz-darwin-amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o build/chatz-darwin-arm64
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o build/chatz-windows-amd64.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="$(LDFLAGS)" -o build/chatz-windows-i386.exe

clean:
	rm -rf chatz build
