#!/bin/sh

cd "$(dirname "$0")/.."

env GOOS=linux GOARCH=amd64 go build -o face-tome ./cmd/face-tome || exit $?

export COMPOSE_FILE=./deployments/docker-compose.yml
docker-compose up app
