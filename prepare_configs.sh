#!/bin/bash

SERVER_HOST=${HELLO_SERVER_HOST:-localhost}
SERVER_PORT=${HELLO_SERVER_PORT:-8080}
SERVER_READ_TO=${HELLO_SERVER_READ_TO:-5s}
SERVER_WRITE_TO=${HELLO_SERVER_WRITE_TO:-10s}
MYSQL_ADDRESS=${HELLO_MYSQL_ADDRESS:-127.0.0.1}
MYSQL_DATABASE=${HELLO_MYSQL_DATABASE:-db_hello}
PROJECT_DIR=$(pwd)

config_files="app.conf /external/hello.nginx.conf"

for filename in $config_files; do
    src=./conf/$filename.tpl
    dest=./conf/$filename
    if test -f "$dest"; then
        echo "File $dest already exist, skipping..."
    else
        sed "s?HELLO_SERVER_HOST?$SERVER_HOST?g" $src | \
        sed "s?HELLO_SERVER_PORT?$SERVER_PORT?g" | \
        sed "s?HELLO_SERVER_READ_TO?$SERVER_READ_TO?g" | \
        sed "s?HELLO_SERVER_WRITE_TO?$SERVER_WRITE_TO?g" | \
        sed "s?HELLO_MYSQL_ADDRESS?$MYSQL_ADDRESS?g" | \
        sed "s?HELLO_MYSQL_DATABASE?$MYSQL_DATABASE?g" | \
        sed "s?HELLO_DIR?$PROJECT_DIR?g" > $dest
        echo $dest created
    fi
done
