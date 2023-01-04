build:
	go build -o social-network-app ./cmd/main.go

dev:
	go run cmd/main.go

swag:
	swag init -g cmd/main.go

migrate:
	migrate -path ./migrations -database 'postgres://postgres:just_code03@0.0.0.0:5432/social?sslmode=disable' up
