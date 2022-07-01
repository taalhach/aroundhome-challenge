FROM golang:alpine as builder
ENV GO111MODULE=on
LABEL maintainer="Muhammad Talha <talhach891@gmail.com>"
RUN apk update && apk add --no-cache git bash make build-base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build
FROM alpine:latest
WORKDIR /root/
COPY configs/app-configs-docker.ini  .
ENV SETTINGS=/root/app-configs-docker.ini
COPY .data/mock_data.xlsx  .
ENV DATA_SOURCE=/root/mock_data.xlsx
COPY --from=builder /app/bin/aroundhome .
RUN apk add bash
EXPOSE 3000
LABEL Name=aroundhome Version=0.0.1
