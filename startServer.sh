#!/bin/bash

echo "Recriando o BD..."
source createTablesMysql.sh

echo "Iniciando o webserver..."
source runGoModTidy.sh
docker exec -it goapp go run ./cmd/walletcore/main.go