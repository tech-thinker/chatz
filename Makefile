VERSION := $(or $(AppVersion), "v1.1.2")
COMMIT := $(or $(shell git rev-parse --short HEAD), "unknown")
BUILDDATE := $(shell date +%Y-%m-%d)

LDFLAGS := -X 'main.AppVersion=$(VERSION)' -X 'main.CommitHash=$(COMMIT)' -X 'main.BuildDate=$(BUILDDATE)'

PLATFORMS = linux darwin windows
ARCHITECTURES = amd64 arm64 arm

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

dist:
	@for platform in $(PLATFORMS); do \
		for arch in $(ARCHITECTURES); do \
			extension=""; if [ "$$platform" = "windows" ]; then extension=".exe"; fi; \
            CGO_ENABLED=0 GOOS=$$platform GOARCH=$$arch go build -ldflags="$(LDFLAGS)" -o build/chatz-$$platform-$$arch$$extension; \
			if [ ! -f build/chatz-$$platform-$$arch ]; then continue; fi; \
			if [ "$$platform" = "windows" ]; then continue; fi; \
			cp build/chatz-$$platform-$$arch build/chatz; \
			tar -zcvf build/chatz-$$platform-$$arch.tar.gz build/chatz man/chatz.1; \
        done \
	done
	rm build/chatz

clean:
	rm -rf chatz build
