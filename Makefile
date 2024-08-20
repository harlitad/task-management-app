run-local: run-swagger
	make build-local
	./main

build-local:
	go build ./cmd/main.go

run-swagger:
	swag init -g cmd/main.go -o docs

run-docker: run-swagger
	docker compose build --no-cache
	docker compose -f docker-compose.yaml up -d

stop-docker:
	docker compose -f docker-compose.yaml down
