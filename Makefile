install:
	go mod tidy

test:
	go test -v ./...  -race -coverprofile=coverage.out -covermode=atomic

run: build
	./chat "This is test."

build:
	go build -o chat .

build-all:
	GOOS=linux GOARCH=amd64 go build -o build/chat-linux-amd64
	GOOS=linux GOARCH=arm64 go build -o build/chat-linux-arm64
	GOOS=linux GOARCH=arm go build -o build/chat-linux-arm
	GOOS=darwin GOARCH=amd64 go build -o build/chat-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -o build/chat-darwin-arm64
	GOOS=windows GOARCH=amd64 go build -o build/chat-windows-amd64.exe
	GOOS=windows GOARCH=386 go build -o build/chat-windows-i386.exe

clean:
	rm -rf chat build
