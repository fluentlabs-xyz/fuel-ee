#!/bin/bash

set -ex

CUR_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd "$CUR_DIR"

docker compose up -d
docker compose stop --timeout 0 redis
docker container rm all_in_one-redis-1
docker volume rm all_in_one_redis_data
docker compose up -d

notify-send "redis: recreation finished"
