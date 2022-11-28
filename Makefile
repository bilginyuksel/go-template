files_to_include_coverage=`go list ./... | grep -v mock | grep -v cmd`
coverage_threshold=80

up:
	docker-compose up

run:
	APP_ENV=local go run ./cmd

lint:
	golangci-lint run

test:
	go test -coverprofile=coverage.out -covermode=atomic ${files_to_include_coverage}

test-unit:
	go test -short ./...

coverage: @test
	chmod +x scripts/code_coverage.sh
	sh scripts/code_coverage.sh

mockgen:
	mockgen -source=internal/expense/service.go -destination=internal/expense/mock/repository.go -package=mock
