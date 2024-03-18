all: run

run:
	docker compose down; docker compose up --build

test:
	go test -cover ./internal/handlers
# 	go test -coverprofile=cover.out ./internal/handlers && go tool cover -html=cover.out -o cover.html
