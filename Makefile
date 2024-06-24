run-local: 
	swag init -g ./cmd/main.go
	make build-local
	./main

build-local:
	go build ./cmd/main.go

run-swagger:
	swag init

run-docker: 
	swag init -g ./cmd/main.go
	docker compose build --no-cache
	docker compose -f docker-compose.yaml up

stop-docker:
	docker compose -f docker-compose.yaml down
