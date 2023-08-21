#! /bin/sh

until etcdctl --endpoints=http://etcd:2379 endpoint health; do
  echo "Waiting..."
  sleep 1
done

CONFIG_CONTENT="$(cat /config/config.yaml)"

etcdctl --endpoints=http://etcd:2379 put /minitok/config "${CONFIG_CONTENT}"
