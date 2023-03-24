local-migrate-up:
	migrate -path ./migrations -database postgres://postgres:postgres@localhost:5432/weather?sslmode=disable up

local-migrate-down:
	migrate -path ./migrations -database postgres://postgres:postgres@localhost:5432/weather?sslmode=disable down

local-docker-up:
	docker compose -f ./deploy/docker-compose.yml up