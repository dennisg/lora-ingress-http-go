#!/usr/bin/env bash

set -e

export PROJECT_ID="things-running"
export APP_NAME="$(basename $PWD)"
export SERVICE_ACCOUNT_NAME="agritech-lora-ingress-${USER}"
export REGION="europe-west3"
export CONCURRENCY="10"

echo "Building: eu.gcr.io/${PROJECT_ID}/${APP_NAME}"
docker build -t eu.gcr.io/${PROJECT_ID}/${APP_NAME} .

echo "Pushing: eu.gcr.io/${PROJECT_ID}/${APP_NAME}"
docker push eu.gcr.io/${PROJECT_ID}/${APP_NAME}

echo "Deploying: $APP_NAME"
gcloud run deploy ${APP_NAME} \
    --project=${PROJECT_ID} \
    --service-account=${SERVICE_ACCOUNT_NAME}@things-running.iam.gserviceaccount.com \
    --allow-unauthenticated \
    --concurrency=${CONCURRENCY} \
    --platform=managed \
    --region=${REGION} \
    --image=eu.gcr.io/${PROJECT_ID}/agritech/${APP_NAME}
