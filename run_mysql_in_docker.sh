#!/usr/bin/env bash

docker run -d \
  --name mysql-employees \
  -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=college \
  -v "${PWD}"/mysql_data:/var/lib/mysql \
  genschsa/mysql-employees
