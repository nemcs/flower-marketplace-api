run:
	go run ./cmd/api

lint:
	golangci-lint run

test:
	go test ./...

docker-up:
	docker-compose -f deploy/docker-compose.yml up -d

docker-down:
	docker-compose -f deploy/docker-compose.yml down
