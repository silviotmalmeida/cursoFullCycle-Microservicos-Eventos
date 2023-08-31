#!/bin/bash

echo "Atualizando as dependências..."
docker exec -it goclientapp go mod tidy
