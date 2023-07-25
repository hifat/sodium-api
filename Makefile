run:
	go run ./cmd/api

migrate:
	go run ./cmd/migration

swag:
	# swag init --parseDependency --parseInternal --output ./docs --generalInfo=./cmd/api/main.go
	swag init --generalInfo=./cmd/api/main.go

rung:
	swag init --generalInfo=./cmd/api/main.go
	go run ./cmd/api/main.go

wire:
	wire ./...
