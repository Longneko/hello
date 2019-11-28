#!/bin/bash

# base app config
DEST=./conf/app.conf
if test -f "$DEST"; then
    echo "File $DEST already exist, skipping..."
else
    cp ./conf/app.conf.tpl ./conf/app.conf
    echo $DEST created
fi

# nginx config
DEST=./conf/external/hello.nginx.conf
if test -f "$DEST"; then
    echo "File $DEST already exist, skipping..."
else
    HOST=${HELLO_SERVER_HOST:-localhost}
    PORT=${HELLO_SERVER_PORT:-8080}
    WORKDIR=$(pwd)
    sed "s?HELLO_SERVER_HOST?$HOST?g" ./conf/external/hello.nginx.conf.tpl | \
    sed "s?HELLO_SERVER_PORT?$PORT?g" | \
    sed "s?HELLO_DIR?$WORKDIR?g" > $DEST
    echo $DEST created
fi
