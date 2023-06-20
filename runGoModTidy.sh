#!/bin/bash

echo "Atualizando as dependências..."
docker exec -it goapp go mod tidy
