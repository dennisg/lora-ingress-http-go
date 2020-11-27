#!/usr/bin/env bash
# used to test the cloud build file that is used for automated deploys

set -e

source .env
export SERVICE_ACCOUNT_NAME="agritech-lora-ingress-${USER}"

echo "Deploying..."

gcloud builds submit \
    --project=${PROJECT_ID} \
    --substitutions="_ACCESS_TOKEN=${ACCESS_TOKEN},_TOPIC=${TARGET_TOPIC},_SERVICE_ACCOUNT_NAME=${SERVICE_ACCOUNT_NAME}" \
    .
