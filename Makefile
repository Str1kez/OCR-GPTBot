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
format:
	gofmt -s -w -l .
run: build
	.bin/bot
