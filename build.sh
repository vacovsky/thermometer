#!/bin/bash

rm greenhouse.zip
GOARCH=arm CGO_ENABLED=1 CC=arm-linux-gnueabihf-gcc GOARM=6 go build -o greenhouse
# zip -r greenhouse.zip greenhouse static bin
zip -r greenhouse.zip ./*
scp -r greenhouse.zip 192.168.111.23:/home/joe/greenhouse/