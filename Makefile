run:
	go run cmd/main.go
watch:
	npx @tailwindcss/cli -i ./src/input.css -o ./src/output.css --watch
dockerenv:
	docker compose --env-file .env-docker up -d --build