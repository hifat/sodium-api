run:
	go run ./cmd/hifatBlogAPI/main.go

migrate:
	go run ./cmd/migration/main.go

swag:
	# swag init --parseDependency --parseInternal --output ./docs --generalInfo=./cmd/hifatBlogAPI/main.go
	swag init --generalInfo=./cmd/hifatBlogAPI/main.go