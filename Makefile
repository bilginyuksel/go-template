run:
	go run .

lint:
	golangci-lint run

test:
	go test ./...

test-unit:
	go test -short ./...
