#!/usr/bin/env bash

set -e

export PORT=8080
export PROJECT_ID="things-running"
export APP_NAME="$(basename $PWD)"
export GOOGLE_APPLICATION_CREDENTIALS=$PWD/things-running-dev-4b06ebd5e307.json
export IMAGE=eu.gcr.io/${PROJECT_ID}/agritech/agritech-lora-ingress

docker build -t ${IMAGE} .

docker run \
   -p 8080:${PORT} \
   -e PORT=${PORT} \
   -e K_SERVICE=dev \
   -e ACCESS_TOKEN=972f36b3-569c-462d-8ca6-9e6b6c8205dd \
   -e K_CONFIGURATION=dev \
   -e K_REVISION=dev-00001 \
   -e GOOGLE_APPLICATION_CREDENTIALS=/tmp/keys/things-running.json \
   -v ${GOOGLE_APPLICATION_CREDENTIALS}:/tmp/keys/things-running.json:ro \
   ${IMAGE}
