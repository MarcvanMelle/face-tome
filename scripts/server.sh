#!/bin/sh

cd "$(dirname "$0")/.."

env GOOS=linux GOARCH=amd64 go build || exit $?

docker-compose up app
