#!/bin/bash


SERVER_HOST=${HELLO_SERVER_HOST:-localhost}
SERVER_PORT=${HELLO_SERVER_PORT:-8080}
SERVER_READ_TO=${HELLO_SERVER_READ_TO:-5s}
SERVER_WRITE_TO=${HELLO_SERVER_WRITE_TO:-10s}
MYSQL_ADDRESS=${HELLO_MYSQL_ADDRESS:-127.0.0.1}
MYSQL_DATABASE=${HELLO_MYSQL_DATABASE:-db_hello}
WORKDIR=$(pwd)

config_files="app.conf /external/hello.nginx.conf"


# for FILENAME in app.conf /external/hello.nginx.conf; do
for FILENAME in $config_files; do
    SRC=./conf/$FILENAME.tpl
    DEST=./conf/$FILENAME
    if test -f "$DEST"; then
        echo "File $DEST already exist, skipping..."
    else
        sed "s?HELLO_SERVER_HOST?$SERVER_HOST?g" $SRC | \
        sed "s?HELLO_SERVER_PORT?$SERVER_PORT?g" | \
        sed "s?HELLO_SERVER_READ_TO?$SERVER_READ_TO?g" | \
        sed "s?HELLO_SERVER_WRITE_TO?$SERVER_WRITE_TO?g" | \
        sed "s?HELLO_MYSQL_ADDRESS?$MYSQL_ADDRESS?g" | \
        sed "s?HELLO_MYSQL_DATABASE?$MYSQL_DATABASE?g" | \
        sed "s?HELLO_DIR?$WORKDIR?g" > $DEST
        echo $DEST created
    fi
done

# # base app config
# DEST=./conf/app.conf
# if test -f "$DEST"; then
#     echo "File $DEST already exist, skipping..."
# else
#     cp ./conf/app.conf.tpl ./conf/app.conf
#     echo $DEST created
# fi
#
# # nginx config
# DEST=./conf/external/hello.nginx.conf
# if test -f "$DEST"; then
#     echo "File $DEST already exist, skipping..."
# else
#     sed "s?HELLO_SERVER_HOST?$SERVER_HOST?g" ./conf/external/hello.nginx.conf.tpl | \
#     sed "s?HELLO_SERVER_PORT?$SERVER_PORT?g" | \
#     sed "s?HELLO_SERVER_READ_TO?$SERVER_READ_TO?g" | \
#     sed "s?HELLO_SERVER_WRITE_TO?$SERVER_WRITE_TO?g" | \
#     sed "s?HELLO_MYSQL_ADDRESS?$MYSQL_ADDRESS?g" | \
#     sed "s?HELLO_MYSQL_DATABASE?$MYSQL_DATABASE?g" | \
#     sed "s?HELLO_DIR?$WORKDIR?g" > $DEST
#     echo $DEST created
# fi
#
