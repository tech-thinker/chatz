install:
	go mod tidy

test:
	go test -v ./...  -race -coverprofile=coverage.out -covermode=atomic

run: build
	./chat "This is test."

build:
	go build -o chat .

build-all:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/chat-linux-amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/chat-linux-arm64
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o build/chat-linux-arm
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/chat-darwin-amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/chat-darwin-arm64
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/chat-windows-amd64.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o build/chat-windows-i386.exe

clean:
	rm -rf chat build
