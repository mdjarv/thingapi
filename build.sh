#!/bin/sh
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o thingapi .
docker build -t thingapi:latest .