.PHONY: up down logs swag

up:
	docker compose up -d --build

down:
	docker compose down

logs:
	docker compose logs -f

swag:
	cd backend &&  swag init -g /cmd/EduCourse/main.go -o ./docs

.DEFAULT_GOAL := up