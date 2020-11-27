#!/usr/bin/env bash

set -e

source .env

export APP_NAME="$(basename $PWD)"
export SERVICE_ACCOUNT_NAME="agritech-lora-ingress-${USER}"
export REGION="europe-west3"
export CONCURRENCY="10"

export IMAGE="eu.gcr.io/${PROJECT_ID}/${ROOT}/${APP_NAME}"
echo "Building: $IMAGE"
docker build -t $IMAGE .

echo "Pushing: $IMAGE"
docker push $IMAGE

echo "Deploying: $APP_NAME"
gcloud run deploy ${APP_NAME} \
    --project=${PROJECT_ID} \
    --service-account=${SERVICE_ACCOUNT_NAME}@${PROJECT_ID}.iam.gserviceaccount.com \
    --allow-unauthenticated \
    --concurrency=${CONCURRENCY} \
    --platform=managed \
    --region=${REGION} \
    --set-env-vars="ACCESS_TOKEN=${ACCESS_TOKEN},TARGET_TOPIC=${TARGET_TOPIC}" \
    --image=${IMAGE}
