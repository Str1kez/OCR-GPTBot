FROM golang:1.20-alpine3.17 AS builder

RUN apk add --no-cache tesseract-ocr-dev tesseract-ocr-data-rus g++

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./.bin/app ./cmd/bot/main.go

FROM alpine:3.17

RUN apk add --no-cache tesseract-ocr-dev tesseract-ocr-data-rus 

WORKDIR /app

COPY --from=builder /app/.bin/app .
COPY ./config ./config

# COPY ./photo_2023-03-13_02-21-14.jpg ./photo_2023-03-13_02-21-14.jpg 

CMD ./app
