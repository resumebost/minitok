#! /usr/bin/env bash

project="minitok"
services=("user" "video" "favorite" "comment")
work_root=$(pwd)

if [[ "$work_root" != */${project} ]]; then
    echo "Should be executed in ${project}"
    exit 1
fi

for service in "${services[@]}"; do
    cd "${work_root}" || exit 1
    kitex -I idl -module ${project} idl/"${service}"/service.proto
    cd "${work_root}"/cmd/"${service}" || exit 1
    kitex -I ../../idl -module ${project} -service "${service}" \
        -use ${project}/kitex_gen \
        ../../idl/"${service}"/service.proto
    rm -f ./kitex_info.yaml
done
