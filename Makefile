run:
	go run ./cmd/api

migrate:
	go run ./cmd/migration

swag:
	# swag init --parseDependency --parseInternal --output ./docs --generalInfo=./cmd/api
	swag init --generalInfo=./cmd/api

rung:
	swag init --generalInfo=./cmd/api
	go run ./cmd/api
