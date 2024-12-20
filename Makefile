VERSION := $(or $(AppVersion), "v0.0.0")
COMMIT := $(or $(shell git rev-parse --short HEAD), "unknown")
BUILDDATE := $(shell date +%Y-%m-%d)

LDFLAGS := -X 'main.AppVersion=$(VERSION)' -X 'main.CommitHash=$(COMMIT)' -X 'main.BuildDate=$(BUILDDATE)'

PLATFORMS = linux darwin windows
ARCHITECTURES = amd64 arm64 arm

all: build

tidy:
	go mod tidy

test:
	go test -v ./...  -race -coverprofile=coverage.out -covermode=atomic

run: build
	./chatz "This is test."

install: build
	@cp man/chatz.1 man/chatz.old
	@sed -e "s|BUILDDATE|$(BUILDDATE)|g" -e "s|VERSION|$(VERSION)|g" man/chatz.old > man/chatz.1
	@cp chatz /usr/local/bin/chatz
	@cp man/chatz.1 /usr/local/share/man/man1/chatz.1
	@mv man/chatz.old man/chatz.1
	@echo "chatz successfully installed."

uninstall:
	@rm /usr/local/bin/chatz
	@rm /usr/local/share/man/man1/chatz.1
	@echo "chatz successfully uninstalled."

build:
	go build -ldflags="$(LDFLAGS)" -o chatz .

dist:
	@cp man/chatz.1 man/chatz.old
	@sed -e "s|BUILDDATE|$(BUILDDATE)|g" -e "s|VERSION|$(VERSION)|g" man/chatz.old > man/chatz.1
	@for platform in $(PLATFORMS); do \
		for arch in $(ARCHITECTURES); do \
			if [ "$$platform" = "darwin" ] && [ "$$arch" = "arm" ]; then continue; fi; \
			extension=""; if [ "$$platform" = "windows" ]; then extension=".exe"; fi; \
            CGO_ENABLED=0 GOOS=$$platform GOARCH=$$arch go build -ldflags="$(LDFLAGS)" -o build/chatz-$$platform-$$arch$$extension; \
			if [ ! -f build/chatz-$$platform-$$arch ]; then continue; fi; \
			if [ "$$platform" = "windows" ]; then continue; fi; \
			cp build/chatz-$$platform-$$arch build/chatz; \
			tar -zcvf build/chatz-$$platform-$$arch.tar.gz build/chatz man/chatz.1; \
        done \
	done
	@rm build/chatz
	@mv man/chatz.old man/chatz.1
	# Generating checksum
	@cd build && sha256sum * > ../checksum-sha256sum.txt
	@cd build && md5sum * > checksum-md5sum.txt
	@cd build && mv ../checksum-sha256sum.txt checksum-sha256sum.txt
	@echo "Checksum generated successfully."

clean:
	rm -rf chatz build
