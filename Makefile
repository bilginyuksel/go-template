run:
	go run ./cmd

lint:
	golangci-lint run

test:
	go test ./...

test-unit:
	go test -short ./...

coverage:
	chmod +x scripts/code_coverage.sh
	sh scripts/code_coverage.sh
