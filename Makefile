.SILENT:

env:
	cp .env.example .env
build:
	go build -o .bin/bot cmd/bot/main.go
format:
	gofmt -s -w -l .
run: build
	.bin/bot

