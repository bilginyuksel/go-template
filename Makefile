run:
	go run .

lint:
	golangci-lint run

test:
	go test ./...

test-unit:
	go test -short ./...

coverage:
	go test -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -func=coverage.out | grep total | awk '{print $3}'
	rm coverage.out
