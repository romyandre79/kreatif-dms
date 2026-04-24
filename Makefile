.PHONY: up down dev-back dev-front sqlc migrate-up migrate-down

up:
	docker-compose up -d

down:
	docker-compose down

test:
	cd backend && go test -v ./...

dev-back:
	cd backend && set APP_ENV=development && go run cmd/server/main.go

staging-back:
	cd backend && set APP_ENV=staging && go run cmd/server/main.go

prod-back:
	cd backend && set APP_ENV=production && go run cmd/server/main.go

dev-front:
	cd frontend && npm run dev

staging-front:
	cd frontend && npm run build:staging

prod-front:
	cd frontend && npm run build:prod

sqlc:
	cd backend && sqlc generate

swagger-gen:
	cd backend && swag init -g cmd/server/main.go

migrate-up:
	cd backend && go run cmd/server/main.go migrate up

migrate-down:
	cd backend && go run cmd/server/main.go migrate down
