#! /usr/bin/env bash

project="minitok"
services=("user" "video" "favorite" "comment")
work_root=$(pwd)

if [[ "$work_root" != */${project} ]]; then
    echo "Should be executed in ${project}"
    exit 1
fi

rm -rf ./kitex_gen

for service in "${services[@]}"; do
    rm -rf ./cmd/${service}/*
done
