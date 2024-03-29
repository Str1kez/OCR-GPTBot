.PHONY: build
.SILENT:

# On Mac
# export LIBRARY_PATH="/opt/homebrew/lib"
# export CPATH="/opt/homebrew/include"

env:
	cp .env.example .env
build:
	go build -o .bin/bot cmd/bot/main.go
build-image:
	docker build . -t str1kez/ocrgpt_bot -f build/Dockerfile
container:
	docker compose -f build/docker-compose.yaml up -d --remove-orphans
db:
	docker compose -f build/docker-compose.yaml up -d --remove-orphans db
format:
	gofumpt -w .
run: build
	MODE=dev .bin/bot
prod-run: build
	MODE=prod .bin/bot
ngrok:
	# for mac only, use --net=host instead host.docker.internal on linux
	docker run --rm -it --env-file .local.env ngrok/ngrok:latest http host.docker.internal:80
