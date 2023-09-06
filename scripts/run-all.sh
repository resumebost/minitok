#! /usr/bin/env bash

project="minitok"
services=("user" "video" "favorite" "comment")
work_root=$(pwd)

if [[ "$work_root" != */${project} ]]; then
    echo "Should be executed in ${project}"
    exit 1
fi

# docker containers
docker-compose up -d

# api service
cd "${work_root}"/cmd/api || exit 1
go run . -x &

for service in "${services[@]}"; do
    cd "${work_root}"/cmd/"${service}" || exit 1
    bash build.sh && output/bootstrap.sh &
done
