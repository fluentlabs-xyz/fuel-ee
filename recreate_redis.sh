#!/bin/bash

set -ex

CUR_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd "$CUR_DIR"

CONTAINER_PREFIX=fuel-ee
docker compose up -d
docker compose stop --timeout 0 redis
docker container rm ${CONTAINER_PREFIX}-redis-1
docker volume rm ${CONTAINER_PREFIX}_redis_data
docker compose up -d
