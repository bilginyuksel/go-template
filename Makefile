up:
	docker-compose up

run:
	APP_ENV=local go run ./cmd

lint:
	golangci-lint run

test:
	go test ./...

test-unit:
	go test -short ./...

coverage:
	chmod +x scripts/code_coverage.sh
	sh scripts/code_coverage.sh

mockgen:
	mockgen -source=internal/expense/service.go -destination=internal/expense/mock/repository.go -package=mock
