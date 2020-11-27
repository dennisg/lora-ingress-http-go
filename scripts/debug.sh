#!/usr/bin/env bash

set -e

source .env

export ENV="dev"

export PORT=8080
export APP_NAME="$(basename $PWD)"
export GOOGLE_APPLICATION_CREDENTIALS=$PWD/${PROJECT_ID}-${ENV}-4b06ebd5e307.json
export IMAGE=eu.gcr.io/${PROJECT_ID}/${ROOT}/${APP_NAME}

docker build -t ${IMAGE} .
docker image ls ${IMAGE}

docker run \
   -p 8080:${PORT} \
   -e PORT=${PORT} \
   -e K_SERVICE=${ENV} \
   -e ACCESS_TOKEN=${ACCESS_TOKEN} \
   -e TARGET_TOPIC=${TARGET_TOPIC} \
   -e K_CONFIGURATION=${ENV} \
   -e K_REVISION=${ENV}-00001 \
   -e GOOGLE_APPLICATION_CREDENTIALS=/tmp/keys/google.json \
   -v ${GOOGLE_APPLICATION_CREDENTIALS}:/tmp/keys/google.json:ro \
   ${IMAGE}
