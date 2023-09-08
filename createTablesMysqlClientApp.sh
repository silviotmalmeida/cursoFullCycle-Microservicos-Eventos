#!/bin/bash

#dados do banco
mysql_user='root'
mysql_password='root'
mysql_database='wallet'

echo "Criando as tabelas..."
docker exec -it mysql-client sh -c "mysql -u ${mysql_user} -p${mysql_password} ${mysql_database} -e 'DROP TABLE IF EXISTS accounts'"
docker exec -it mysql-client sh -c "mysql -u ${mysql_user} -p${mysql_password} ${mysql_database} -e 'CREATE TABLE accounts (id VARCHAR(255), balance INT, PRIMARY KEY (id))'"