#!/bin/bash

#dados do banco
mysql_user='root'
mysql_password='root'
mysql_database='wallet'

echo "Criando as tabelas..."
docker exec -it mysql sh -c "mysql -u ${mysql_user} -p${mysql_password} ${mysql_database} -e 'DROP TABLE IF EXISTS clients'"
docker exec -it mysql sh -c "mysql -u ${mysql_user} -p${mysql_password} ${mysql_database} -e 'DROP TABLE IF EXISTS accounts'"
docker exec -it mysql sh -c "mysql -u ${mysql_user} -p${mysql_password} ${mysql_database} -e 'DROP TABLE IF EXISTS transactions'"
docker exec -it mysql sh -c "mysql -u ${mysql_user} -p${mysql_password} ${mysql_database} -e 'CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at DATE, updated_at DATE)'"
docker exec -it mysql sh -c "mysql -u ${mysql_user} -p${mysql_password} ${mysql_database} -e 'CREATE TABLE accounts (id VARCHAR(255), client_id VARCHAR(255), balance INT, created_at DATE)'"
docker exec -it mysql sh -c "mysql -u ${mysql_user} -p${mysql_password} ${mysql_database} -e 'CREATE TABLE transactions (id VARCHAR(255), account_id_from VARCHAR(255), account_id_to VARCHAR(255), amount INT, created_at DATE)'"
