lint:
	golangci-lint run ./...

run: lint
	go run ./cmd

test: lint
	go test -v -count=1 -race -gcflags=-l -timeout=30s ./...

build: test
	go build -o ./build/mail-service -ldflags "-s -w" ./cmd

build-linux: test
	GOOS=linux GOARCH=amd64 go build -o ./build/mail-service -ldflags "-s -w" ./cmd