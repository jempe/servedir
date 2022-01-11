#!/bin/bash

CONFIG_FOLDER=~/.servedir

if [-d $CONFIG_FOLDER];
then
	echo "$CONFIG_FOLDER already exists"
else
	mkdir $CONFIG_FOLDER
fi

openssl genrsa -out $CONFIG_FOLDER/server.key 2048

openssl req -new -x509 -sha256 -key $CONFIG_FOLDER/server.key -out $CONFIG_FOLDER/server.crt -days 3650

chmod 0700 $CONFIG_FOLDER/server.crt
chmod 0700 $CONFIG_FOLDER/server.key

