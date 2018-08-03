#!/bin/sh

cd "$(dirname "$0")/../.."

export COMPOSE_FILE=./face-tome/deployments/docker-compose.yml
docker-compose run test go test -timeout 5000ms ./... -v -race
