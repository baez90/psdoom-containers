#!/usr/bin/env bash

for (( i = 0; i < 20; ++i )); do
    docker run --rm -d --name "alpine_$i" alpine:latest /bin/ash -c 'sleep 3600';
done
