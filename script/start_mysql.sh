#!/bin/bash

set -eux

cd $(dirname $0)/..
wd=$(pwd)

docker-compose up -d

sleep 10

mysql -u root -h 127.0.0.1 -p root -P 3306 sample < ./docker/mysql/init/create-tables.sql
