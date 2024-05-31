#! /usr/bin/env bash
cd `dirname $0`

# [[ $1 == '' ]] && echo "Usage: $BASH_SOURCE <container_id>" && exit 1
# CONTAINER=$1

CONTAINER=`docker ps | awk '/mysql-snippetbox/{print $1}'`
# echo $CONTAINER && exit 0

source ../../env/mysql.env
echo Creating tables...
docker exec -i $CONTAINER sh -c "exec mysql -u$MYSQL_ROOT_USER -p$MYSQL_ROOT_PASSWORD" < ./create_tables.sql
echo Insert data...
docker exec -i $CONTAINER sh -c "exec mysql -u$MYSQL_ROOT_USER -p$MYSQL_ROOT_PASSWORD" < ./insert_data.sql
echo Creating users...
docker exec -i $CONTAINER sh -c "exec mysql -u$MYSQL_ROOT_USER -p$MYSQL_ROOT_PASSWORD" < ./create_users.sql