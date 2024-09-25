VERSION := $(or $(AppVersion), "v1.1.2")
COMMIT := $(or $(shell git rev-parse --short HEAD), "unknown")
BUILDDATE := $(shell date +%Y-%m-%d)

LDFLAGS := -X 'main.AppVersion=$(VERSION)' -X 'main.CommitHash=$(COMMIT)' -X 'main.BuildDate=$(BUILDDATE)'

all: build

dep:
	go mod tidy

test:
	go test -v ./...  -race -coverprofile=coverage.out -covermode=atomic

run: build
	./chatz "This is test."

install: build
	cp chatz /usr/local/bin/chatz
	cp man/chatz.1 /usr/local/share/man/man1/chatz.1

build:
	go build -ldflags="$(LDFLAGS)" -o chatz .

build-all:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o build/chatz-linux-amd64
	cp build/chatz-linux-amd64 build/chatz
	tar -zcvf build/chatz-linux-amd64.tar.gz build/chatz man/chatz.1
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o build/chatz-linux-arm64
	cp build/chatz-linux-arm64 build/chatz
	tar -zcvf build/chatz-linux-arm64.tar.gz build/chatz man/chatz.1
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="$(LDFLAGS)" -o build/chatz-linux-arm
	cp build/chatz-linux-arm build/chatz
	tar -zcvf build/chatz-linux-arm.tar.gz build/chatz man/chatz.1
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o build/chatz-darwin-amd64
	cp build/chatz-darwin-amd64 build/chatz
	tar -zcvf build/chatz-darwin-amd64.tar.gz build/chatz man/chatz.1
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o build/chatz-darwin-arm64
	cp build/chatz-darwin-arm64 build/chatz
	tar -zcvf build/chatz-darwin-arm64.tar.gz build/chatz man/chatz.1
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o build/chatz-windows-amd64.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="$(LDFLAGS)" -o build/chatz-windows-i386.exe
	rm build/chatz

clean:
	rm -rf chatz build
