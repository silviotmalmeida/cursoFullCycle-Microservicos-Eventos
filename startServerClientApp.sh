#!/bin/bash

echo "Recriando o BD..."
source createTablesMysqlClientApp.sh

echo "Iniciando o webserver..."
source runGoModTidyClientApp.sh
docker exec -it goclientapp go run ./cmd/main.go